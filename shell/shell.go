package shell

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type Session struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
	DryRun bool
}

func NewSession(in io.Reader, out, errs io.Writer) *Session {
	return &Session{Stdin: in, Stdout: out, Stderr: errs}
}

func CmdFromString(input string) (*exec.Cmd, error) {
	args := strings.Fields(input)
	if len(args) < 1 {
		return &exec.Cmd{}, fmt.Errorf("empty input")
	}
	return exec.Command(args[0], args[1:]...), nil
}

func (s *Session) Run() {
	fmt.Fprint(s.Stdout, "> ")
	input := bufio.NewScanner(s.Stdin)
	for input.Scan() {
		line := input.Text()
		cmd, err := CmdFromString(line)
		if err != nil {
            fmt.Fprintf(s.Stdout, "> ")
			continue
		}
        if s.DryRun {
            fmt.Fprintf(s.Stdout, "%s\n> ", line)
            continue
        }
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintln(s.Stderr, "err", err)
		}
		fmt.Fprintf(s.Stdout, "%s", out)
	}
    fmt.Fprintln(s.Stdout, "\nSee you later!")
}

func Main() {
    session := NewSession(os.Stdin, os.Stdout, os.Stderr)
    session.Run()
}