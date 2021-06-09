package qmsg

import (
	"errors"

	"github.com/saitofun/qlib/net/qbuf"
)

type Parser interface {
	Marshal(qbuf.Buffer, Message) error
	Unmarshal(qbuf.Buffer) (Message, error)
}

var ErrParseTCPDataLack = errors.New("data lack")
