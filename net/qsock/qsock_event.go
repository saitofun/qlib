package qsock

import (
	"git.querycap.com/ss/lib/net/qsock/qmsg"
)

type Event struct {
	node    *Node
	payload qmsg.Message
}

func (ev *Event) From() *Node {
	return ev.node
}

func (ev *Event) Payload() qmsg.Message {
	return ev.payload
}

func (ev *Event) Response(rsp qmsg.Message) error {
	return ev.node.WriteMessage(rsp)
}

func (ev *Event) Send(msg qmsg.Message) error {
	return ev.node.SendMessage(msg)
}

// for trace only
func (ev *Event) Endpoint() *Node { return ev.node }
