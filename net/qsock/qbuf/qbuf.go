package qbuf

import "errors"

type Buffer interface {
	Read([]byte) (int, error)
	Write([]byte) (int, error)
	ResetAndWrite([]byte) error
	Probe(int) ([]byte, error)
	Shift(int) error
	Bytes() []byte
	Reset()
	State() (int, int, int) // read size/write size/capacity
}

var EStreamBufferDataLack = errors.New("stream buffer date lack")
var EPacketBufferDataLack = errors.New("packet buffer data lack")
