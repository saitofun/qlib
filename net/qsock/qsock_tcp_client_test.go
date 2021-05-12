package qsock_test

import (
	"fmt"
	"testing"
	"time"

	"git.querycap.com/ss/lib/encoding/qjson"
	"git.querycap.com/ss/lib/net/qsock"
	"git.querycap.com/ss/lib/net/qsock/qmsg"
)

func TestTCPClient(t *testing.T) {
	var req = NewGreetingReq("hello")
	var cli, err = qsock.NewClient(
		qsock.ClientOptionParser(TCPParser),
		qsock.ClientOptionRemote("localhost:10086"),
		qsock.ClientOptionProtocol(qsock.ProtocolTCP),
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

	req.Content = "hello"
	for {
		fmt.Printf("-> %s\n", qjson.UnsafeMarshalString(req))
		rsp, err := cli.Request(req)
		if err != nil {
			if qsock.IsNodeClosedError(err) {
				break
			}
			fmt.Println(err)
			time.Sleep(3 * time.Second)
			continue
		}
		fmt.Printf("<- %s, %d\n", qjson.UnsafeMarshalString(rsp),
			(rsp.(qmsg.WithTimestamp).GetTimestamp()-req.GetTimestamp())/1e6)
		time.Sleep(250 * time.Millisecond)
		req.Renew()
	}
}
