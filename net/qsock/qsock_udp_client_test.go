package qsock_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/saitofun/qlib/encoding/qjson"
	"github.com/saitofun/qlib/net/qsock"
	"github.com/saitofun/qlib/net/qsock/qmsg"
)

func TestUDPClient(t *testing.T) {
	var req = NewGreetingReq("first greeting req")
	var cli, err = qsock.NewClient(
		qsock.ClientOptionParser(Parser),
		qsock.ClientOptionRemote("localhost:10010"),
		qsock.ClientOptionProtocol(qsock.ProtocolUDP),
		qsock.ClientOptionNodeID("udp_client"),
		qsock.ClientOptionTimeout(time.Second),
		qsock.ClientOptionOnConnected(func(n *qsock.Node) {
			req.Content = "first greeting req"
			for {
				req.Renew()
				rsp, err := n.Request(req, 2*time.Second)
				if err != nil {
					fmt.Printf("cli: OnConnected error:%v\n", err)
					if qsock.IsTimeoutError(err) {
						continue
					}
					if qsock.IsNodeClosedError(err) {
						break
					}
				}
				msg, ok := rsp.(*GreetingRsp)
				if ok && msg.Content == "first greeting rsp" {
					fmt.Println("OnConnected done")
					break
				}
			}
		}),
		qsock.ClientOptionRoute(rsp.Type(), func(ev *qsock.Event) {
			fmt.Printf("route -> %s\n", qjson.UnsafeMarshal(ev.Payload()))
		}),
	)
	if err != nil {
		panic(err)
	}

	for {
		if cli.IsClosed() {
			break
		}
		fmt.Printf("-> %s\n", qjson.UnsafeMarshalString(req))
		rsp, err := cli.Request(req)
		if err != nil {
			if cli.IsClosed() {
				break
			}
			fmt.Println(err)
			time.Sleep(3 * time.Second)
			continue
		}
		fmt.Printf("<- %s, %d\n", qjson.UnsafeMarshalString(rsp),
			rsp.(qmsg.WithTimestamp).GetTimestamp()-req.GetTimestamp())
		time.Sleep(time.Millisecond * 250)
		req.Renew()
	}
	fmt.Println("client closed")
}
