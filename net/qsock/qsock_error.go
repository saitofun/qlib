package qsock

import (
	"fmt"
)

type Errno int

const (
	ENodeTimeout Errno = 0x01 + iota
	ENodeContextDone
	ENodeMessage // received nil message,
	ENodeMarshal
	ENodeUnmarshal
	ENodeClosed
	ENodeDial
	ENodeRead
	ENodeWrite
	ENodeListen
	ENodeResolve
	ENodeInvalidParser
	ENodeInvalidRemoteAddr
	ENodeInvalidProtocol
	ENodeInvalidListenAddr
	ENodeOption
	EMessageUnbound
	EMessageTimeout
)

var es = [...]string{
	ENodeTimeout:           "qsock.Timeout",
	ENodeContextDone:       "qsock.Context",
	ENodeMessage:           "qsock.Message",
	ENodeMarshal:           "qsock.Marshal",
	ENodeUnmarshal:         "qsock.Unmarshal",
	ENodeClosed:            "qsock.Closed",
	ENodeDial:              "qsock.Dial",
	ENodeRead:              "qsock.Read",
	ENodeWrite:             "qsock.Write",
	ENodeListen:            "qsock.Listen",
	ENodeResolve:           "qsock.Resolve",
	ENodeInvalidParser:     "qsock.InvalidParser",
	ENodeInvalidRemoteAddr: "qsock.InvalidRemote",
	ENodeInvalidProtocol:   "qsock.InvalidProtocol",
	ENodeInvalidListenAddr: "qsock.InvalidListen",
	ENodeOption:            "qsock.Option",
	EMessageUnbound:        "qsock.UnboundMessage",
	EMessageTimeout:        "qsock.MessageTimeout",
}

func (e Errno) Error() string {
	if e >= ENodeTimeout && e <= EMessageTimeout {
		return es[e]
	}
	return ""
}

func (e Errno) WithMessage(msg string) *ErrnoWithMsg {
	return &ErrnoWithMsg{
		Errno: e,
		msg:   msg,
	}
}

func (e Errno) WithError(err error) *ErrnoWithErr {
	return &ErrnoWithErr{
		Errno: e,
		err:   err,
	}
}

type ErrnoWithMsg struct {
	Errno
	msg string
}

func (e *ErrnoWithMsg) Error() string {
	return fmt.Sprintf("%v: %s", e.Errno, e.msg)
}

func (e *ErrnoWithMsg) Unwrap() error {
	return e.Errno
}

type ErrnoWithErr struct {
	Errno
	err error
}

func (e *ErrnoWithErr) Error() string {
	return fmt.Sprintf("%v: %s", e.Errno, e.err.Error())
}

func (e *ErrnoWithErr) Unwrap() error {
	return e.Errno
}

func IsTimeoutError(e error) bool {
	switch v := e.(type) {
	case Errno:
		return v == EMessageTimeout || v == ENodeTimeout
	case *ErrnoWithMsg:
		return v.Errno == EMessageTimeout || v.Errno == ENodeTimeout
	case *ErrnoWithErr:
		return v.Errno == EMessageTimeout || v.Errno == ENodeTimeout
	default:
		return false
	}
}

func IsNodeClosedError(e error) bool {
	switch v := e.(type) {
	case Errno:
		return v == ENodeClosed
	case *ErrnoWithMsg:
		return v.Errno == ENodeClosed
	case *ErrnoWithErr:
		return v.Errno == ENodeClosed
	default:
		return false
	}
}
