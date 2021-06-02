package qipc

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/saitofun/qlib/net/qsock"
	"github.com/saitofun/qlib/net/qsock/qmsg"
)

type ipc struct {
	pid         int
	sock        string
	srv         *qsock.Server
	srvInstOnce *sync.Once
	clients     map[int]*qsock.Client
	clientsMtx  *sync.Mutex

	*option
}

func (v *ipc) client(peer int) (*qsock.Client, error) {
	v.clientsMtx.Lock()
	defer v.clientsMtx.Unlock()
	if c, ok := v.clients[peer]; ok {
		return c, nil
	} else {
		c, err := qsock.NewClient(
			qsock.ClientOptionParser(v.parser),
			qsock.ClientOptionReadBufferCap(v.bufferSize),
			qsock.ClientOptionWriteBufferCap(v.bufferSize),
			qsock.ClientOptionRemote(fmt.Sprintf("%s%d.sock", TMPDIR, peer)),
			qsock.ClientOptionTimeout(v.timeout),
			qsock.ClientOptionProtocol(qsock.ProtocolUnix),
			qsock.ClientOptionNodeID(strconv.Itoa(peer)),
		)
		if err != nil {
			return nil, err
		}
		v.clients[peer] = c
		return c, nil
	}
}

func (v *ipc) Request(pid int, req IPCMessage) (IPCMessage, error) {
	var (
		cli, err = v.client(pid)
		rsp      qmsg.Message
	)
	if err != nil {
		return nil, err
	}
	rsp, err = cli.Request(req)
	if err != nil {
		return nil, err
	}
	if rsp == nil {
		return nil, nil
	}
	if v, ok := rsp.(IPCMessage); ok {
		return v, nil
	}
	return nil, nil
}

func (v *ipc) ReceiveEvent() (*qsock.Event, error) {
	var err error
	if v.parser == nil {
		return nil, qsock.ENodeInvalidParser
	}

	if v.srv == nil {
		v.srvInstOnce.Do(func() {
			v.srv, err = qsock.NewServer(
				qsock.ServerOptionTimeout(v.timeout),
				qsock.ServerOptionReadBufferCap(v.bufferSize),
				qsock.ServerOptionWriteBufferCap(v.bufferSize),
				qsock.ServerOptionListenAddr(v.sock),
				qsock.ServerOptionProtocol(qsock.ProtocolUnix),
				qsock.ServerOptionParser(v.parser),
			)
		})
		if err != nil {
			return nil, err
		}
	}
	return v.srv.ReceiveEvent(), nil
}

func (v *ipc) Pid() int { return v.pid }

var v *ipc

const TMPDIR = "/tmp/qproc/"

func init() {
	var pid = os.Getpid()
	var err = os.MkdirAll(TMPDIR, 0777)
	if err != nil {
		panic(err)
	}
	v = &ipc{
		pid:         pid,
		sock:        fmt.Sprintf("%s%d.sock", TMPDIR, pid),
		srv:         nil,
		srvInstOnce: &sync.Once{},
		clients:     make(map[int]*qsock.Client),
		clientsMtx:  &sync.Mutex{},
		option: &option{
			bufferSize: 32 * 1024,
			parser:     nil,
			timeout:    time.Second << 1,
		},
	}
}

func SetMessageParser(parser qmsg.Parser) {
	v.parser = parser
}

func ReceiveEvent() (*IPCEvent, error) {
	ev, err := v.ReceiveEvent()
	if err != nil {
		return nil, err
	}
	return &IPCEvent{ev.Endpoint(), ev.Payload().(IPCMessage)}, nil
}
func Request(pid int, req IPCMessage) (IPCMessage, error) {
	return v.Request(pid, req)
}
