package qproc

import "git.querycap.com/aisys/lib/qsock"

type IPCEvent struct {
	node *qsock.Node
	msg  IPCMessage
}

func (v *IPCEvent) Payload() IPCMessage {
	return v.msg
}

func (v *IPCEvent) Endpoint() *qsock.Node {
	return v.node
}

func (v *IPCEvent) Response(msg IPCMessage) error {
	return v.node.WriteMessage(msg)
}

func (v *IPCEvent) Send(msg IPCMessage) error {
	return v.node.SendMessage(msg)
}