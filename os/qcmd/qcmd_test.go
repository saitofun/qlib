package qcmd

import (
	"fmt"
	"testing"
)

func TestCmd(t *testing.T) {
	var cmd = New()
	// err = cmd.Exec("ls", "-al", "/Users/sincos")
	// if err == nil {
	// 	for _, l := range cmd.Lines() {
	// 		fmt.Println(l)
	// 	}
	// }

	// err = cmd.Exec("uname -s; uname -r; uname -v")
	// if err == nil {
	// 	for _, l := range cmd.Lines() {
	// 		fmt.Println(l)
	// 	}
	// 	fmt.Println("*")
	// }

	ls, err := cmd.Exec("cat /Users/sincos/main.go | grep 'wtf'")
	if err == nil {
		fmt.Println(ls)
		return
	}
	fmt.Println(err)
}
