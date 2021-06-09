package qipc

import (
	"github.com/saitofun/qlib/net/qmsg"
)

type IPCMessage interface {
	SrcPid() int
	DstPid() int
	qmsg.Message
}
