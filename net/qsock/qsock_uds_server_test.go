package qsock_test

import (
	"fmt"
	"testing"

	"git.querycap.com/ss/lib/net/qsock"
)

func TestUDSServer(t *testing.T) {
	var addr = "/tmp/ipc.sock"
	var srv, err = qsock.NewServer(
		qsock.ServerOptionConnCap(10),
		qsock.ServerOptionListenAddr(addr),
		qsock.ServerOptionProtocol(qsock.ProtocolUnix),
		qsock.ServerOptionParser(Parser),
		qsock.ServerOptionRoute(typ, HandleGreeting))

	fmt.Println("server started: " + addr)

	if err != nil {
		panic(err)
	}
	srv.Serve()
}
