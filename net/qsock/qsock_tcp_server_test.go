package qsock_test

import (
	"fmt"
	"testing"

	"github.com/saitofun/qlib/net/qsock"
)

func TestTCPServer(t *testing.T) {
	var addr = ":10086"
	var srv, err = qsock.NewServer(
		qsock.ServerOptionConnCap(10),
		qsock.ServerOptionListenAddr(addr),
		qsock.ServerOptionProtocol(qsock.ProtocolTCP),
		qsock.ServerOptionParser(TCPParser),
		qsock.ServerOptionRoute(typ, HandleGreeting),
	)
	// qsock.ServerOptionOnConnected(func(n *qsock.Node) {
	// 	for {
	// 		msg, err := n.ReadMessage(time.Second)
	// 		if err != nil {
	// 			fmt.Printf("srv: OnConnected [error:%v]\n", err)
	// 			if qsock.IsTimeoutError(err) {
	// 				continue
	// 			}
	// 			if qsock.IsNodeClosedError(err) {
	// 				break
	// 			}
	// 		}
	// 		raw, ok := msg.(*GreetingReq)
	// 		if ok && raw.Content == "first greeting req" {
	// 			err = n.WriteMessage(&GreetingRsp{
	// 				Header: Header{
	// 					Id:        raw.Id,
	// 					Typ:       "greeting_rsp",
	// 					Timestamp: time.Now().UnixNano(),
	// 				},
	// 				Content: "first greeting rsp",
	// 				Pid:     os.Getegid(),
	// 			})
	// 			if err != nil {
	// 				fmt.Printf("srv: OnConnected [error:%v]\n", err)
	// 				break
	// 			}
	// 			fmt.Println("srv: OnConnected done")
	// 			break
	// 		}
	// 		continue
	// 	}
	// }))

	fmt.Println("server started: " + addr)

	if err != nil {
		panic(err)
	}
	srv.Serve()
}
