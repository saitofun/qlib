package qsock

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/saitofun/qlib/net/qbuf"
	"github.com/saitofun/qlib/net/qbuf/qbuf_packet"
	"github.com/saitofun/qlib/net/qbuf/qbuf_stream"
	"github.com/saitofun/qlib/net/qmsg"
	"github.com/saitofun/qlib/os/qsche"
)

type Server struct {
	*ServerOption
	mgr    *clients
	seq    int
	ln     net.Listener
	udp    *net.UDPConn
	buf    qbuf.Buffer
	worker qsche.WorkersScheduler
}

func NewServer(options ...ServerOptionSetter) (*Server, error) {
	var srv = &Server{
		mgr: NewClientManager(),
		seq: 1,
		ServerOption: &ServerOption{
			timeout:        5 * time.Second,
			readBufferCap:  DefaultBufferCapacity,
			writeBufferCap: DefaultBufferCapacity,
			workerPoolSize: 1024,
		},
	}

	for _, opt := range options {
		opt(srv.ServerOption)
	}
	if srv.listenAddr == "" {
		return nil, ENodeInvalidListenAddr
	}
	if srv.parser == nil {
		return nil, ENodeInvalidParser
	}
	if srv.protocol == ProtocolUnknown {
		return nil, ENodeInvalidProtocol
	}
	if srv.routes != nil || srv.handler != nil {
		// TODO concurrency option
		srv.worker = qsche.RunConScheduler(srv.workerPoolSize)
	}

	err := srv.listen()
	if err != nil {
		return nil, err
	}

	go srv.run()

	return srv, nil
}

func (s *Server) listen() (err error) {
	if s.protocol == ProtocolUnix {
		os.Remove(s.listenAddr)
	}
	if s.protocol == ProtocolUDP {
		var addr *net.UDPAddr
		addr, err = net.ResolveUDPAddr(s.protocol.Network(), s.listenAddr)
		if err != nil {
			return ENodeResolve.WithError(err)
		}
		s.udp, err = net.ListenUDP(s.protocol.Network(), addr)
		if err == nil {
			err = SetConnOption(s.udp, s.readBufferCap, s.writeBufferCap)
		}
	} else {
		s.ln, err = net.Listen(s.protocol.Network(), s.listenAddr)
	}
	if err != nil {
		return ENodeListen.WithError(err)
	}
	return
}

func (s *Server) buffers() (qbuf.Buffer, qbuf.Buffer) {
	if s.protocol == ProtocolTCP {
		return qbuf_stream.New(s.readBufferCap), qbuf_packet.New(s.writeBufferCap)
	}
	return qbuf_packet.New(s.readBufferCap), qbuf_packet.New(s.writeBufferCap)
}

func (s *Server) income(conn net.Conn, addr net.Addr) *Node {
	n := NewNode()
	n.rq = make(chan qmsg.Message, 1024)
	n.sq = make(chan qmsg.Message, 1024)
	n.rb, n.wb = s.buffers()
	n.parser = s.parser
	n.protocol = s.protocol
	n.c = conn
	if addr != nil {
		n.id = addr.String()
		n.addr = addr
	} else {
		n.id = strconv.Itoa(s.seq)
		s.seq++
	}
	n.Run()
	if n.protocol != ProtocolUDP {
		for _, f := range s.onConnected {
			f(n)
		}
	}
	s.mgr.New(n, s.onDisconnected...)
	return n
}

func (s *Server) run() {
	var err error
	if s.ln != nil {
		var conn net.Conn
		for {
			conn, err = s.ln.Accept()
			if err != nil {
				fmt.Printf("listener: %v\n", err)
				continue
			}
			err = SetConnOption(conn, s.readBufferCap, s.writeBufferCap)
			if err != nil {
				fmt.Printf("option: %v\n", err)
				continue
			}
			n := s.income(conn, nil)
			n.id = fmt.Sprintf("[%s - %s - %s]",
				s.ln.Addr().Network(), s.ln.Addr(), conn.RemoteAddr())
			fmt.Printf("client: %v connected\n", n.id)
		}
	} else {
		var (
			buf  = make([]byte, DefaultBufferCapacity)
			n    *Node
			addr *net.UDPAddr
			msg  qmsg.Message
			size = 0
		)

		for {
			if err != nil {
				fmt.Printf("listener: %v\n", err)
			}
			buf = buf[0:cap(buf)]
			size, addr, err = s.udp.ReadFromUDP(buf)
			if err != nil {
				continue
			}
			n = s.mgr.Get(addr.String())
			if n == nil {
				n = s.income(s.udp, addr)
				fmt.Printf("client: %v newNode created\n", n.id)
			}
			if err = n.rb.ResetAndWrite(buf[0:size]); err != nil {
				continue
			}
			if msg, err = n.parser.Unmarshal(n.rb); err != nil {
				continue
			}
			n.rq <- msg
		}
	}
}

func (s *Server) Clients() *clients { return s.mgr }

func (s *Server) ReceiveEvent() *Event { return <-s.mgr.events }

func (s *Server) Serve() {
	if s.routes == nil && s.handler == nil {
		return
	}
	for {
		ev := s.ReceiveEvent()
		if handlers := s.routes.Handlers(ev.payload.Type()); len(handlers) > 0 {
			for _, h := range handlers {
				s.worker.Add(HandlerFunc(h, ev))
			}
			continue
		}
		if s.handler != nil {
			s.worker.Add(HandlerFunc(s.handler, ev))
		}
	}
}
