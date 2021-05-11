package qipc

import (
	"git.querycap.com/ss/lib/net/qsock/qmsg"
)

type IPCMessage interface {
	SrcPid() int
	DstPid() int
	qmsg.Message
}
