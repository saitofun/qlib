package qbuf_stream

import (
	"testing"
)

func TestStreamBuffer(t *testing.T) {
	var (
		buf  = New(50)
		dat  = make([]byte, 15)
		left = buf.Cap()
	)

	t.Logf("r=%3d w=%3d rs=%3d ws=%3d\n", buf.r, buf.w, buf.rs, buf.ws)
	for left > len(dat) {
		buf.Write(dat)
		left -= len(dat)
		t.Logf("r=%3d w=%3d rs=%3d ws=%3d\n", buf.r, buf.w, buf.rs, buf.ws)
	}

	buf.Read(dat)
	t.Logf("r=%3d w=%3d rs=%3d ws=%3d AfterRead15\n", buf.r, buf.w, buf.rs, buf.ws)

	buf.Write(dat)
	t.Logf("r=%3d w=%3d rs=%3d ws=%3d AfterWrite15\n", buf.r, buf.w, buf.rs, buf.ws)

	if _, err := buf.Probe(buf.rs); err != nil {
		t.Error(err)
	}
	sft := buf.rs - 2
	buf.Shift(sft)
	t.Logf("r=%3d w=%3d rs=%3d ws=%3d AfterShift%d\n", buf.r, buf.w, buf.rs, buf.ws, sft)

	dat = make([]byte, 50)
	buf.Write(dat)
	t.Logf("r=%3d w=%3d rs=%3d ws=%3d AfterShift50 cap=%d\n", buf.r, buf.w, buf.rs, buf.ws, buf.Cap())
}
