package qbuf

import "errors"

type Buffer interface {
	Len() int                     // Len buffer read length
	Cap() int                     // Cap buffer capacity
	Read(out []byte) (int, error) // Read read len(out)
	Write(in []byte) (int, error) // Write write in
	ResetAndWrite([]byte) error   // ResetAndWrite reset buffer and write
	Probe(int) ([]byte, error)    // Probe read without shift read position
	Shift(int) error              // Shift shift read position
	Bytes() []byte
	Reset()
}

var EStreamBufferDataLack = errors.New("stream buffer date lack")
var EPacketBufferDataLack = errors.New("packet buffer data lack")
