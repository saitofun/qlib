package qtype

import (
	"git.querycap.com/ss/lib/container/qlist"
	"git.querycap.com/ss/lib/encoding"
)

type pool struct {
	busy map[int]*qlist.List
	free map[int]*qlist.List
}

var p = &pool{
	busy: make(map[int]*qlist.List),
	free: make(map[int]*qlist.List),
}

func (p *pool) get(grd int) []byte {
	if grd == -1 {
		return nil
	}
	if v := p.free[grd].PopFront(); v != nil {
		return v.([]byte)[0:0]
	}
	v := make([]byte, 0, 2<<grd)
	p.busy[grd].PushBack(v)
	return v
}

func (p *pool) put(dat []byte) {
	if cap(dat) == 0 {
		return
	}
	p.free[grade(cap(dat))].PushBack(dat)
}

func grade(n int) (pow int) {
	return 0
}

type Bytes struct {
	buf []byte
	grd int
}

func NewBytes() *Bytes {
	return &Bytes{grd: -1}
}

func NewBytesString(s string) *Bytes {
	grd := grade(len(s))
	return (&Bytes{buf: p.get(grd), grd: grd}).
		AppendString(s)
}

func NewBytesCap(cap int) *Bytes {
	grd := grade(cap)
	return &Bytes{buf: p.get(grd), grd: grd}
}

func (b *Bytes) Val() []byte {
	return b.buf
}

func (b *Bytes) Clone() *Bytes {
	ret := &Bytes{}
	ret.buf = p.get(b.grd)
	copy(ret.buf, b.buf)
	ret.grd = b.grd
	return ret
}

func (b *Bytes) Append(bytes ...byte) *Bytes {
	b.try(len(bytes))
	b.buf = append(b.buf, bytes...)
	return b
}

func (b *Bytes) AppendString(s string) *Bytes {
	b.try(len(s))
	b.buf = append(b.buf, encoding.StrToBytes(s)...)
	return b
}

func (b *Bytes) AppendBytes(v *Bytes) *Bytes {
	return b.Append(v.buf...)
}

func (b *Bytes) Len() int { return len(b.buf) }

func (b *Bytes) Cap() int { return cap(b.buf) }

func (b *Bytes) try(delta int) {
	if len(b.buf)+delta > cap(b.buf) {
		buf := p.get(b.grd + 1)
		buf = append(buf, b.buf...)
		p.put(b.buf)
		b.buf = buf
		b.grd++
	}
	return
}

func (b *Bytes) QuoteBy(v byte) *Bytes {
	var buf []byte
	if len(b.buf)+2 > cap(b.buf) {
		buf = p.get(b.grd + 1)
	} else {
		buf = p.get(b.grd)
	}
	buf = append(buf, v)
	buf = append(buf, b.buf...)
	buf = append(buf, v)
	p.put(b.buf)
	b.buf = buf
	return b
}

func ContactBytes(bytes ...*Bytes) *Bytes {
	total := 0
	for i := range bytes {
		total += len(bytes[i].buf)
	}
	grd := grade(total)
	ret := &Bytes{
		buf: p.get(grd),
		grd: grd,
	}
	for i := range bytes {
		copy(ret.buf, bytes[i].buf)
	}
	return ret
}

func ReleaseBytes(bytes ...*Bytes) {
	for i := range bytes {
		p.put(bytes[i].buf)
	}
}
