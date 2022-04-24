package qlist_test

import (
	"fmt"

	"github.com/saitofun/qlib/container/qlist"
)

var (
	List     *qlist.List
	SafeList *qlist.List
	capacity = 10
)

func New() {
	l := qlist.New()

	for i := 0; i < capacity; i++ {
		l.PushBack(i)
	}

	fmt.Println(l.Len())
	fmt.Println(l)
	fmt.Println(l.String())

	List = l
}

func NewSafe() {
	l := qlist.NewSafe()

	for i := 0; i < capacity; i++ {
		l.PushBack(i)
	}

	fmt.Println(l.Len())
	fmt.Println(l)
	fmt.Println(l.String())

	SafeList = l
}

func ExampleNew() {
	New()
	NewSafe()

	// Output:
	// 10
	// [0,1,2,3,4,5,6,7,8,9]
	// [0,1,2,3,4,5,6,7,8,9]
	// 10
	// [0,1,2,3,4,5,6,7,8,9]
	// [0,1,2,3,4,5,6,7,8,9]
}
