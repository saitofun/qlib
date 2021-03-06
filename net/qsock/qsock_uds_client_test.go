package qsock_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/saitofun/qlib/encoding/qjson"
	"github.com/saitofun/qlib/net/qmsg"
	"github.com/saitofun/qlib/net/qsock"
)

func TestUDSClient(t *testing.T) {
	var cli, err = qsock.NewClient(
		qsock.ClientOptionParser(Parser),
		qsock.ClientOptionRemote("/tmp/ipc.sock"),
		qsock.ClientOptionProtocol(qsock.ProtocolUnix),
	)
	if err != nil {
		panic(err)
	}

	req := NewGreetingReq("hello")
	for {
		fmt.Printf("-> %s\n", qjson.UnsafeMarshalString(req))
		rsp, err := cli.Request(req)
		if err != nil {
			fmt.Println(err)
			time.Sleep(3 * time.Second)
			continue
		}
		fmt.Printf("<- %s, %d\n", qjson.UnsafeMarshalString(rsp),
			rsp.(qmsg.WithTimestamp).GetTimestamp()-req.GetTimestamp())
		time.Sleep(time.Second)
		req.Renew()
	}
}
