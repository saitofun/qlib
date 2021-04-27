// Package qtype qtype.Bytes provide global byte slice
package qtype

import (
	"runtime"
	"sync"

	"git.querycap.com/ss/lib/container/qlist"
)

// 支持0～4M []byte的分配关联, 所有大于等于4M的[]byte保留在最后一个槽
var (
	gGrade    = 22 + 1
	gElements = 4
)

type pool struct {
	caps      []int
	busy      []*qlist.List
	free      []*qlist.List
	clr       []chan struct{}
	permanent map[string]bytes
	mu        *sync.Mutex
}

var p *pool

func init() {
	p = &pool{
		caps: make([]int, gGrade),
		busy: make([]*qlist.List, gGrade),
		free: make([]*qlist.List, gGrade),
		clr:  make([]chan struct{}, gGrade),
		mu:   &sync.Mutex{},
	}
	for i := 1; i < gGrade; i++ {
		p.caps[i] = 1 << i
		p.busy[i] = qlist.NewSafe()
		p.free[i] = qlist.NewSafe(make([]byte, 0, p.caps[i]))
		p.clr[i] = make(chan struct{}, 10)
	}
	return
}

func (p *pool) get(need int) *qlist.Element {
	if need == 0 {
		return nil
	}

	n := grade(need)

	p.mu.Lock()
	defer p.mu.Unlock()

	for {
		if buf := p.free[n].PopFront(); buf == nil {
			p.free[n].PushBack(make([]byte, 0, p.caps[n]))
		} else {
			return p.busy[n].PushBack(buf)
		}
	}
}

func (p *pool) put(ele *qlist.Element) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if ele == nil {
		return
	}

	val := ele.Value.(bytes)
	if n := grade(cap(val)); n != 0 {
		p.busy[n].Remove(ele)
		p.free[n].PushBack(val)
	}
}

func grade(len int) (grd int) {
	for (1 << grd) < len {
		grd++
	}
	if grd > 21 {
		grd = 21
	}
	return
}

type bytes = []byte
type Bytes struct {
	dat bytes
	ele *qlist.Element
}

func NewBytes() (ret *Bytes) {
	ret = &Bytes{}
	defer runtime.SetFinalizer(ret.ele, p.put)
	return ret
}

func NewBytesString(s string) *Bytes {
	return &Bytes{}
}

func NewPermanent(cap int) *Bytes {
	return &Bytes{}
}

func (b *Bytes) Val() []byte {
	return nil
}

func (b *Bytes) Clone() *Bytes {
	return b
}

func (b *Bytes) Append(bytes ...byte) *Bytes {
	return b
}

func (b *Bytes) AppendString(s string) *Bytes {
	return b
}

func (b *Bytes) AppendBytes(bytes ...*Bytes) *Bytes {
	return b
}

func (b *Bytes) Len() int { return len(b.dat) }

func (b *Bytes) Cap() int { return cap(b.dat) }

func (b *Bytes) String() string { return string(b.dat) }

func (b *Bytes) try(delta int) {
	return
}

func ContactBytes(bytes ...*Bytes) *Bytes {
	return nil
}

func ReleaseBytes(bytes ...*Bytes) {
	return
}

type BytesPoolStatus struct {
	Free    [][2]int `json:"free"`
	FreeSum int      `json:"free_sum"`
	Busy    [][2]int `json:"busy"`
	BusySum int      `json:"busy_sum"`
	Sum     int      `json:"sum"`
}

func GetBytesPoolStatus() *BytesPoolStatus {
	ret := &BytesPoolStatus{}

	p.mu.Lock()
	defer p.mu.Unlock()

	for i := 1; i < len(p.free); i++ {
		ret.Free = append(ret.Free, [2]int{p.caps[i], p.free[i].Len()})
		p.free[i].Range(func(e *qlist.Element) bool {
			ret.FreeSum += cap(e.Value.([]byte))
			return true
		})
	}
	for i := 1; i < len(p.busy); i++ {
		ret.Busy = append(ret.Busy, [2]int{p.caps[i], p.busy[i].Len()})
		p.busy[i].Range(func(e *qlist.Element) bool {
			ret.BusySum += cap(e.Value.([]byte))
			return true
		})
	}
	ret.Sum = ret.FreeSum + ret.BusySum
	return ret
}
