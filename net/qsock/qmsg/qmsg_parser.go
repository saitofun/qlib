package qmsg

import (
	"errors"

	"git.querycap.com/ss/lib/net/qsock/qbuf"
)

type Parser interface {
	Marshal(qbuf.Buffer, Message) error
	Unmarshal(qbuf.Buffer) (Message, error)
}

var ErrParseTCPDataLack = errors.New("data lack")
