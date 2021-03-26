package qsock

import (
	"net"
	"time"

	"git.querycap.com/aisys/lib/net/qsock/qmsg"
)

type ProtocolType uint8

const (
	ProtocolUnknown ProtocolType = iota
	ProtocolTCP
	ProtocolUDP
	ProtocolUnix
	ProtocolQUIC // @todo
)

func (p ProtocolType) Network() string {
	switch p {
	case ProtocolTCP:
		return "tcp"
	case ProtocolUDP:
		return "udp"
	case ProtocolUnix:
		return "unix"
	case ProtocolQUIC:
		return "quic"
	default:
		return ""
	}
}

type ClientOption struct {
	remote          string        // remote addr
	protocol        ProtocolType  // protocol type
	timeout         time.Duration // timeout rea/write timeout
	parser          qmsg.Parser   // payload data parser for data conversation
	readBufferCap   int           // user space read buffer capacity
	writeBufferCap  int           // user space write buffer capacity
	readBufferSize  int           // socket option for receive buffer size
	writeBufferSize int           // socket option for write buffer size
	nodeID          string        // client identifier
	workerPoolSize  int           // limited concurrent worker
	onConnected     []func(*Node) // udp ignored
	routes          *Routes       // bind message type and handler
	handler         Handler       // common handler to invoke business API
}

type ClientOptionSetter func(*ClientOption)

func ClientOptionRemote(v string) ClientOptionSetter {
	return func(o *ClientOption) {
		o.remote = v
	}
}

func ClientOptionProtocol(v ProtocolType) ClientOptionSetter {
	return func(o *ClientOption) {
		o.protocol = v
	}
}

func ClientOptionTimeout(v time.Duration) ClientOptionSetter {
	return func(o *ClientOption) {
		o.timeout = v
	}
}

func ClientOptionReadBufferCap(v int) ClientOptionSetter {
	return func(o *ClientOption) {
		o.readBufferCap = v
	}
}

func ClientOptionWriteBufferCap(v int) ClientOptionSetter {
	return func(o *ClientOption) {
		o.writeBufferCap = v
	}
}

func ClientOptionParser(v qmsg.Parser) ClientOptionSetter {
	return func(o *ClientOption) {
		o.parser = v
	}
}

func ClientOptionNodeID(v string) ClientOptionSetter {
	return func(o *ClientOption) {
		o.nodeID = v
	}
}

func ClientOptionReadBufferSize(v int) ClientOptionSetter {
	return func(o *ClientOption) {
		o.readBufferSize = v
	}
}

func ClientOptionWriteBufferSize(v int) ClientOptionSetter {
	return func(o *ClientOption) {
		o.writeBufferSize = v
	}
}

func ClientOptionWorkerPoolSize(v int) ClientOptionSetter {
	return func(o *ClientOption) {
		o.workerPoolSize = v
	}
}

func ClientOptionOnConnected(f func(*Node)) ClientOptionSetter {
	return func(o *ClientOption) {
		o.onConnected = append(o.onConnected, f)
	}
}

func ClientOptionRoute(t qmsg.Type, h ...Handler) ClientOptionSetter {
	return func(o *ClientOption) {
		if o.routes == nil {
			o.routes = NewRoutes()
		}
		o.routes.Register(t, h...)
	}
}

func ClientOptionHandler(h Handler) ClientOptionSetter {
	return func(o *ClientOption) {
		o.handler = h
	}
}

type ServerOption struct {
	listenAddr      string
	protocol        ProtocolType
	timeout         time.Duration
	parser          qmsg.Parser
	connCap         int
	readBufferCap   int
	writeBufferCap  int
	readBufferSize  int
	writeBufferSize int
	onConnected     []func(*Node) // udp ignored
	workerPoolSize  int
	routes          *Routes
	handler         Handler
}

type ServerOptionSetter func(*ServerOption)

func ServerOptionListenAddr(v string) ServerOptionSetter {
	return func(o *ServerOption) {
		o.listenAddr = v
	}
}

func ServerOptionProtocol(v ProtocolType) ServerOptionSetter {
	return func(o *ServerOption) {
		o.protocol = v
	}
}

func ServerOptionTimeout(v time.Duration) ServerOptionSetter {
	return func(o *ServerOption) {
		o.timeout = v
	}
}

func ServerOptionParser(v qmsg.Parser) ServerOptionSetter {
	return func(o *ServerOption) {
		o.parser = v
	}
}

func ServerOptionConnCap(v int) ServerOptionSetter {
	return func(o *ServerOption) {
		o.connCap = v
	}
}

func ServerOptionReadBufferCap(v int) ServerOptionSetter {
	return func(o *ServerOption) {
		o.readBufferCap = v
	}
}

func ServerOptionWriteBufferCap(v int) ServerOptionSetter {
	return func(o *ServerOption) {
		o.writeBufferCap = v
	}
}

func ServerOptionOnConnected(f func(*Node)) ServerOptionSetter {
	return func(o *ServerOption) {
		o.onConnected = append(o.onConnected, f)
	}
}

func ServerOptionReadBufferSize(v int) ServerOptionSetter {
	return func(o *ServerOption) {
		o.readBufferSize = v
	}
}

func ServerOptionWriteBufferSize(v int) ServerOptionSetter {
	return func(o *ServerOption) {
		o.writeBufferSize = v
	}
}

func ServerOptionWorkerPoolSize(v int) ServerOptionSetter {
	return func(o *ServerOption) {
		o.workerPoolSize = v
	}
}

func ServerOptionRoute(t qmsg.Type, h ...Handler) ServerOptionSetter {
	return func(o *ServerOption) {
		if o.routes == nil {
			o.routes = NewRoutes()
		}
		o.routes.Register(t, h...)
	}
}

func ServerOptionHandler(h Handler) ServerOptionSetter {
	return func(o *ServerOption) {
		o.handler = h
	}
}

const (
	DefaultBufferCapacity = 1024 * 1024
)

func SetConnOption(c net.Conn, r, w int) (err error) {
	switch v := c.(type) {
	case *net.TCPConn:
		if err = v.SetReadBuffer(r); err != nil {
			return
		}
		if err = v.SetWriteBuffer(w); err != nil {
			return
		}
		if err = v.SetKeepAlive(true); err != nil {
			return
		}
		if err = v.SetNoDelay(true); err != nil {
			return
		}
	case *net.UDPConn:
		if err = v.SetReadBuffer(r); err != nil {
			return
		}
		if err = v.SetWriteBuffer(w); err != nil {
			return
		}
	case *net.UnixConn:
		if err = v.SetReadBuffer(r); err != nil {
			return
		}
		if err = v.SetWriteBuffer(w); err != nil {
			return
		}
	}
	return nil
}
