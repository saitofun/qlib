package qpool

import (
	"container/list"
	"errors"
	"time"

	"git.querycap.com/ss/lib/container/qlist"
	"git.querycap.com/ss/lib/container/qtype"
	"git.querycap.com/ss/lib/os/qtime"
)

type Pool struct {
	lst    *qlist.List
	closed *qtype.Bool
	*option
}

type item struct {
	val interface{} // val value
	exp int64       // exp expired at
}

type option struct {
	ttl time.Duration
	new func() (interface{}, error)
	exp func(interface{})
}

type PoolOptionSetter func(*option)

func PoolItemTTL(ttl time.Duration) PoolOptionSetter {
	return func(o *option) { o.ttl = ttl }
}

func PoolItemNewFn(fn func() (interface{}, error)) PoolOptionSetter {
	return func(o *option) { o.new = fn }
}

func PoolItemExpFn(fn func(interface{})) PoolOptionSetter {
	return func(o *option) { o.exp = fn }
}

func New(setters ...PoolOptionSetter) *Pool {
	opt := &option{}
	for i := range setters {
		setters[i](opt)
	}
	if opt.new == nil {
		panic("should assign NewFn for pool items")
	}
	// defer timer for check
	return &Pool{lst: qlist.NewSafe(), closed: qtype.NewBool(), option: opt}
}

func (p *Pool) Put(v interface{}) error {
	if p.closed.Val() {
		return errors.New("closed pool")
	}
	item := &item{val: v}
	if p.ttl != 0 {
		item.exp = qtime.NowMillionSecond() + p.ttl.Nanoseconds()/1000000
	}
	p.lst.PushBack(item)
	return nil
}

func (p *Pool) Get() (interface{}, error) {
	for !p.closed.Val() {
		elem, _ := p.lst.PopFront().(*item)
		if elem != nil {
			if elem.exp == 0 || elem.exp > qtime.NowMillionSecond() {
				return elem.val, nil
			} else {
				if p.exp != nil {
					p.exp(elem.val)
				}
			}
		} else {
			break
		}
	}
	if p.new != nil {
		return p.new()
	}
	return nil, errors.New("empty pool")
}

func (p *Pool) Clr() {
	p.lst.LockDo(func(list *list.List) {
		e := list.Front()
		if p.exp != nil {
			for e != nil {
				list.Remove(e)
				p.exp(e.Value.(*item).val)
				e = list.Front()
			}
		} else {
			for e != nil {
				list.Remove(e)
				e = list.Front()
			}
		}
	})
}

func (p *Pool) Len() int { return p.lst.Len() }

func (p *Pool) Close() { p.closed.Set(true) }

func (p *Pool) chk() {
	if p.closed.Val() {
		// pop all elements and exit timer
	}
	if p.ttl == 0 {
		return
	}
}
