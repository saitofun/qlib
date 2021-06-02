package main

import (
	"fmt"

	"github.com/saitofun/qlib/os/qipc"
	"github.com/saitofun/qlib/os/qipc/qipc_test"
)

func main() {
	qipc.SetMessageParser(qipc_test.Parser)
	for {
		ev, err := qipc.ReceiveEvent()
		if err != nil {
			fmt.Println(err)
			continue
		}
		req := ev.Payload()
		fmt.Println("->", req)
		rsp := qipc_test.NewRsp()
		rsp.Id = req.ID().(qipc_test.ID)
		rsp.Msg = "ack"
		ev.Response(rsp)
		fmt.Println("<-", rsp)
	}
}
