package main

import (
	"fmt"
	"time"

	"git.querycap.com/ss/lib/qproc"
	"git.querycap.com/ss/lib/qproc/qproc_test"
	"github.com/google/uuid"
)

func main() {
	qproc.SetMessageParser(qproc_demo.Parser)
	req := qproc_demo.NewReq()
	req.Msg = "syn"
	for {
		req.Id = qproc_demo.ID(uuid.New().String())

		fmt.Println("<-", req)

		rsp, err := qproc.Request(75087, qproc_demo.NewReq())
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Second)
		fmt.Println("->", rsp)
	}
}
