package qsock

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"git.querycap.com/ss/lib/net/qsock/qbuf"
	"git.querycap.com/ss/lib/net/qsock/qmsg"
	"git.querycap.com/ss/lib/qroutines"
)

type NodeTask struct {
	msg qmsg.Message
	src *Node
	dst []net.Addr
}

type Node struct {
	c       net.Conn
	addr    net.Addr
	rb      qbuf.Buffer
	wb      qbuf.Buffer
	rq      chan qmsg.Message
	sq      chan qmsg.Message
	binder  *Binder
	routes  *Routes
	handler Handler
	parser  qmsg.Parser
	worker  *qroutines.LimitedWorkerPool

	id        string
	ctx       context.Context
	cancel    context.CancelFunc
	onceClose *sync.Once
	writerMtx *sync.Mutex
	mgr       *clients
	protocol  ProtocolType
	closed    bool
	reason    string
}

func NewNode() *Node {
	ret := &Node{
		onceClose: &sync.Once{},
		writerMtx: &sync.Mutex{},
		binder:    NewBinder(),
	}
	ret.ctx, ret.cancel = context.WithCancel(context.Background())
	return ret
}

func (n *Node) ID() string {
	return n.id
}

func (n *Node) Context() context.Context {
	return n.ctx
}

func (n *Node) SetQueues(r, s chan qmsg.Message) {
	n.rq, n.sq = r, s
}

func (n *Node) SetParser(parser qmsg.Parser) {
	n.parser = parser
}

func (n *Node) SetBuffers(r, w qbuf.Buffer) {
	n.rb, n.wb = r, w
}

// ReadMessage read message from receive channel
func (n *Node) ReadMessage(d time.Duration) (msg qmsg.Message, err error) {
	if n.closed {
		return nil, ENodeClosed
	}
	if d == 0 {
		return <-n.rq, nil
	}
	select {
	case msg = <-n.rq:
		return msg, nil
	case <-time.After(d):
		return nil, ENodeTimeout
	}
}

// WriteMessage write message to send buffer
func (n *Node) WriteMessage(msg qmsg.Message) (err error) {
	if n.closed {
		return ENodeClosed
	}
	n.writerMtx.Lock()
	defer n.writerMtx.Unlock()
	n.wb.Reset()
	err = n.parser.Marshal(n.wb, msg)
	if err != nil {
		return ENodeMarshal.WithError(err)
	}
	var dat = n.wb.Bytes()
	_, err = n.write(dat)
	if err != nil {
		err = ENodeWrite.WithError(err)
		n.Stop(err)
	}
	return err
}

// SendMessage send message to send queue
func (n *Node) SendMessage(msg qmsg.Message) (err error) {
	if n.closed {
		return ENodeClosed
	}
	defer func() {
		e, ok := recover().(error)
		if ok && e != nil {
			err = e
		}
	}()
	n.sq <- msg
	return nil
}

// Request request message to peer until timeout or responded
func (n *Node) Request(req qmsg.Message, d time.Duration) (rsp qmsg.Message, err error) {
	if n.closed {
		return nil, ENodeClosed
	}
	n.binder.New(req.ID())
	if err = n.WriteMessage(req); err != nil {
		n.binder.Remove(req.ID())
		return nil, err
	}
	return n.binder.Wait(req.ID(), d)
}

// Stop stop Node running
func (n *Node) Stop(reason ...interface{}) {
	n.onceClose.Do(func() {
		if n.cancel != nil {
			n.cancel()
		}
		addr := ""
		if n.c != nil && n.addr == nil {
			addr = n.c.RemoteAddr().String()
			_ = n.c.Close()
		}
		if n.addr != nil {
			addr = n.addr.String()
		}
		if n.sq != nil {
			close(n.sq)
		}
		if n.rq != nil {
			close(n.rq)
		}
		if n.wb != nil {
			n.wb.Reset()
		}
		if n.rb != nil {
			n.rb.Reset()
		}
		if n.binder != nil {
			n.binder.Reset()
		}
		n.closed = true
		n.reason = fmt.Sprint(reason...)
		fmt.Printf("%s cloesd: %v remote: %s\n", n.id, n.reason, addr)
	})
}

// StopReason return stop reason
func (n *Node) StopReason() string { return n.reason }

// Run start Node writing and reading
func (n *Node) Run() {
	go n.recv()
	go n.send()
}

func (n *Node) RunUDP() {
	go n.send()
}

func (n *Node) send() {
	var (
		msg qmsg.Message
		err error
	)

	defer func() {
		n.Stop(err)
	}()

	for {
		select {
		case msg = <-n.sq:
		case <-n.ctx.Done():
			err = ENodeContextDone
			return
		}
		if msg == nil {
			err = ENodeMessage
			return
		}
		if err = n.WriteMessage(msg); err != nil {
			e := errors.Unwrap(err)
			if errors.Is(e, ENodeWrite) {
				return
			}
		}
	}
}

func (n *Node) recv() {
	var (
		msg  qmsg.Message
		err  error
		dat  = make([]byte, DefaultBufferCapacity)
		jobs []qroutines.Job
		size int
	)

	defer func() {
		n.Stop(err)
	}()

	for {
		select {
		case <-n.ctx.Done():
			err = n.ctx.Err()
			return
		default:
			size, err = n.read(dat)
			if err != nil {
				err = ENodeRead.WithError(err)
				return
			}
			if size == 0 {
				err = ENodeRead.WithMessage("read size=0")
				return
			}

			err = n.rb.ResetAndWrite(dat[0:size])
			dat = dat[0:cap(dat)]
			for {
				msg, err = n.parser.Unmarshal(n.rb)
				if err != nil {
					if err == qbuf.EStreamBufferDataLack {
						break
					}
					return
				}
				if n.binder.Push(msg) {
					goto check
				}
				if (n.routes != nil || n.handler != nil) && n.worker != nil {
					if n.routes != nil {
						if jobs = n.routes.GetJobs(&Event{n, msg}); len(jobs) > 0 {
							n.worker.Add(jobs...)
						}
					}
					if n.handler != nil {
						n.worker.Add(func() { n.handler(&Event{n, msg}) })
					}
				}
				n.rq <- msg
			check:
				if n.protocol == ProtocolTCP {
					continue
				}
				break
			}
		}
	}
}

func (n *Node) read(dat []byte) (size int, err error) {
	if n.addr == nil {
		size, err = n.c.Read(dat)
	} else {
		size, _, err = n.c.(*net.UDPConn).ReadFromUDP(dat)
	}
	return
}

func (n *Node) write(dat []byte) (size int, err error) {
	if n.addr != nil {
		size, err = n.c.(*net.UDPConn).WriteTo(dat, n.addr)
		return
	} else {
		written := 0
		for size < len(dat) {
			written, err = n.c.Write(dat[size:])
			if err != nil {
				return
			}
			size += written
		}
	}
	return
}

type NodeStat struct {
	RBufLen int `json:"rb_len"`
	RBufCap int `json:"rb_cap"`
	WBufLen int `json:"wb_len"`
	WBufCap int `json:"wb_cap"`
	RQLen   int `json:"rq_len"`
	SQLen   int `json:"sq_len"`
}

func (n *Node) State() NodeStat {
	return NodeStat{
		n.rb.Len(),
		n.rb.Cap(),
		n.wb.Len(),
		n.wb.Cap(),
		len(n.rq),
		len(n.sq),
	}
}
