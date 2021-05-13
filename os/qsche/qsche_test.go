package qsche

import "testing"

func TestError(t *testing.T) {
	v := &Result{
		Val:   nil,
		error: nil,
	}
	err := func() error { return v }()
	err2 := func() error { return v.error }()
	t.Log(err == nil)
	t.Log(err.(error) == nil)
	t.Log(err2 == nil)
}
