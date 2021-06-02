package qipc

import (
	"time"

	"github.com/saitofun/qlib/net/qsock/qmsg"
)

type option struct {
	bufferSize int
	parser     qmsg.Parser
	timeout    time.Duration
}

type OptionSetter func(*option)

func OptionBufferSize(v int) OptionSetter {
	return func(o *option) {
		o.bufferSize = v
	}
}

func OptionParser(v qmsg.Parser) OptionSetter {
	return func(o *option) {
		o.parser = v
	}
}

func OptionTimeout(v time.Duration) OptionSetter {
	return func(o *option) {
		o.timeout = v
	}
}
