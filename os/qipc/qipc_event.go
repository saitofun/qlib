package qipc

import "git.querycap.com/ss/lib/net/qsock"

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
