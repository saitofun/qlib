package qlist

import (
	"container/list"

	"git.querycap.com/ss/lib/internal/mutex"
)

type Element = list.Element
type List struct {
	mu mutex.Mutex
	list.List
}

func New(values ...interface{}) *List { return nil }

func (l *List) Len() int                                        { return 0 }
func (l *List) PushFront(v interface{}) *Element                { return nil }
func (l *List) PushFrontN(v ...interface{})                     {}
func (l *List) PushFrontList(*List)                             {}
func (l *List) PushBack(v interface{}) *Element                 { return nil }
func (l *List) PushBackN(v ...interface{})                      {}
func (l *List) PushBackList(*List)                              {}
func (l *List) PopFront() interface{}                           { return nil }
func (l *List) PopFrontN() []interface{}                        { return nil }
func (l *List) PopBack() interface{}                            { return nil }
func (l *List) PopBackN() []interface{}                         { return nil }
func (l *List) Front() *Element                                 { return nil }
func (l *List) Back() *Element                                  { return nil }
func (l *List) FrontValue() interface{}                         { return nil }
func (l *List) BackValue() interface{}                          { return nil }
func (l *List) MoveToFront(e *Element)                          {}
func (l *List) MoveToBack(e *Element)                           {}
func (l *List) InsertBefore(v interface{}, e *Element) *Element { return nil }
func (l *List) InsertAfter(interface{}, *Element) *Element      { return nil }
func (l *List) MoveBefore(*Element, *Element)                   {}
func (l *List) MoveAfter(*Element, *Element)                    {}
func (l *List) Remove(...*Element)                              {}
func (l *List) Clear()                                          {}
func (l *List) LockFunc()                                       {}
