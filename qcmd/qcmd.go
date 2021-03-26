package qcmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"strings"
	"syscall"

	"git.querycap.com/aisys/lib/encoding"
)

type Cmd struct {
}

func New() *Cmd {
	return &Cmd{}
}

func (c *Cmd) Exec(args ...string) (ls []string, err error) {
	var (
		arg = make([]string, 2, len(args)+2)
		cmd *exec.Cmd
		pip io.ReadCloser
		out []byte
	)
	arg[0], arg[1] = "sh", "-c"
	arg = append(arg, strings.Join(args, " "))

	cmd = exec.Command(arg[0], arg[1:]...)
	cmd.SysProcAttr = &syscall.SysProcAttr{}

	defer func() {
		if err != nil {
			err = fmt.Errorf("err=%v args=%v lines=%v", err, args, ls)
		}
	}()

	pip, err = cmd.StdoutPipe()
	if err != nil {
		return
	}
	defer pip.Close()

	err = cmd.Start()
	if err != nil {
		return
	}

	out, err = ioutil.ReadAll(pip)
	if err != nil {
		return
	}

	if err = cmd.Wait(); err != nil {
		return nil, err
	}
	if cmd.ProcessState != nil {
		if code := cmd.ProcessState.ExitCode(); code != 0 {
			err = fmt.Errorf("exit_code_%d", code)
			return
		}
	}

	if len(out) > 0 {
		ls = strings.Split(encoding.BytesToStr(out), "\n")
		if out[len(out)-1] == '\n' {
			ls = ls[0 : len(ls)-1]
		}
	}

	return
}
