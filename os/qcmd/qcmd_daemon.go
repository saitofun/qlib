package qcmd

import (
	"os/exec"
	"strings"
)

type Daemon struct {
	cmd *exec.Cmd
	err error
}

func NewDaemonCmd() *Daemon { return &Daemon{} }

func (d *Daemon) Launch(args ...string) {
	var arg = make([]string, 2, len(args)+2)
	arg[0], arg[1] = "sh", "-c"
	arg = append(arg, strings.Join(args, " "))

	d.cmd = exec.Command(arg[0], arg[1:]...)
	go func() {
		d.err = d.cmd.Start()
		if d.err != nil {
			return
		}
		d.err = d.cmd.Wait()
	}()
}

func (d *Daemon) Err() error { return d.err }

// ExitCode return exit code and exited
func (d *Daemon) ExitCode() (int, bool) {
	if d.cmd != nil && d.cmd.ProcessState != nil {
		return d.cmd.ProcessState.ExitCode(), true
	}
	return 0, false
}
