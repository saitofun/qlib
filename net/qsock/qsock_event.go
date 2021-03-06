package qsock

import (
	"github.com/saitofun/qlib/net/qmsg"
)

type Event struct {
	node    *Node
	payload qmsg.Message
}

func NewEvent(node *Node, pld qmsg.Message) *Event {
	return &Event{node, pld}
}

func (ev *Event) Node() *Node { return ev.node }

func (ev *Event) Payload() qmsg.Message { return ev.payload }

func (ev *Event) Response(rsp qmsg.Message) error {
	return ev.node.WriteMessage(rsp)
}

func (ev *Event) Send(msg qmsg.Message) error {
	return ev.node.SendMessage(msg)
}
