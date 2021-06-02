package main

import (
	"fmt"
	"time"

	"github.com/saitofun/qlib/os/qipc"
	"github.com/saitofun/qlib/os/qipc/qipc_test"
	"github.com/google/uuid"
)

func main() {
	qipc.SetMessageParser(qipc_test.Parser)
	req := qipc_test.NewReq()
	req.Msg = "syn"
	for {
		req.Id = qipc_test.ID(uuid.New().String())

		fmt.Println("<-", req)

		rsp, err := qipc.Request(75087, qipc_test.NewReq())
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Second)
		fmt.Println("->", rsp)
	}
}
