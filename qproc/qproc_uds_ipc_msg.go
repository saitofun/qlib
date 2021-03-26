package qproc

import (
	"git.querycap.com/aisys/lib/qsock/qmsg"
)

type IPCMessage interface {
	SrcPid() int
	DstPid() int
	qmsg.Message
}

