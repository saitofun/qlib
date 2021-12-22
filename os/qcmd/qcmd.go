package qcmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/saitofun/qlib/encoding"
)

type Cmd struct {
	cmd *exec.Cmd
	err error
}

func New() *Cmd {
	return &Cmd{}
}

func (c *Cmd) String() string {
	if c.cmd != nil {
		return c.cmd.String()
	}
	return ""
}

func (c *Cmd) Exec(args ...string) (ls []string, err error) {
	var (
		arg = make([]string, 2, len(args)+2)
		pip io.ReadCloser
		out []byte
	)
	arg[0], arg[1] = "sh", "-c"
	arg = append(arg, strings.Join(args, " "))

	c.cmd = exec.Command(arg[0], arg[1:]...)
	c.cmd.SysProcAttr = &syscall.SysProcAttr{}

	defer func() {
		if err != nil {
			err = fmt.Errorf("err=%v args=%v lines=%v", err, args, ls)
			c.err = err
		}
	}()

	pip, err = c.cmd.StdoutPipe()
	if err != nil {
		return
	}
	defer pip.Close()

	err = c.cmd.Start()
	if err != nil {
		return
	}

	out, err = ioutil.ReadAll(pip)
	if err != nil {
		return
	}

	if err = c.cmd.Wait(); err != nil {
		return nil, err
	}
	if c.cmd.ProcessState != nil {
		if code := c.cmd.ProcessState.ExitCode(); code != 0 {
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

func (c *Cmd) Launch(args ...string) {
	c.cmd = exec.Command("sh", "-c", strings.Join(args, " "))
	go func() {
		c.err = c.cmd.Start()
		if c.err != nil {
			return
		}
		c.err = c.cmd.Wait()
	}()
	for {
		time.Sleep(time.Second)
		if c.cmd.Process != nil {
			return
		}
	}
}

func (c *Cmd) Stop() {
	if c.cmd == nil {
		return
	}
	_, exited := c.State()
	if exited {
		return
	}
	if c.cmd.Process != nil {
		_ = syscall.Kill(c.cmd.Process.Pid, syscall.SIGKILL)
	}
	return
}

// State return command exit code and exited
func (c *Cmd) State() (error, bool) {
	if c.err != nil {
		return c.err, true
	}
	if c.cmd != nil && c.cmd.ProcessState != nil {
		return fmt.Errorf("EXIT: %d", c.cmd.ProcessState.ExitCode()), true
	}
	return nil, false
}

func (c *Cmd) Kill() error {
	if c.cmd != nil && c.cmd.Process != nil {
		return c.cmd.Process.Kill()
	}
	return nil
}

func (c *Cmd) Pid() int {
	if c.cmd != nil && c.cmd.Process != nil {
		return c.cmd.Process.Pid
	}
	return 0
}
