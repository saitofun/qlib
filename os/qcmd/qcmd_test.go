package qcmd

import (
	"testing"
)

func TestCmd(t *testing.T) {
	var cmd = New()
	cmd.Launch("sleep 1000000")
	t.Log(cmd.State())
	t.Log(cmd.Pid())
}
