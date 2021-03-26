package main

import (
	"fmt"

	"git.querycap.com/aisys/lib/qproc"
	"git.querycap.com/aisys/lib/qproc/qproc_test"
)

func main() {
	qproc.SetMessageParser(qproc_demo.Parser)
	for {
		ev, err := qproc.ReceiveEvent()
		if err != nil {
			fmt.Println(err)
			continue
		}
		req := ev.Payload()
		fmt.Println("->", req)
		rsp := qproc_demo.NewRsp()
		rsp.Id = req.ID().(qproc_demo.ID)
		rsp.Msg = "ack"
		ev.Response(rsp)
		fmt.Println("<-", rsp)
	}
}
