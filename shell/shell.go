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
    Transcript io.Writer
	DryRun bool
}

func NewSession(in io.Reader, out, errs io.Writer) *Session {
	return &Session{Stdin: in, Stdout: out, Stderr: errs, Transcript: io.Discard}
}

func CmdFromString(input string) (*exec.Cmd, error) {
	args := strings.Fields(input)
	if len(args) < 1 {
		return nil, fmt.Errorf("empty input")
	}
	return exec.Command(args[0], args[1:]...), nil
}

func (s *Session) Run() {
    stdout := io.MultiWriter(s.Stdout, s.Transcript)
    stderr := io.MultiWriter(s.Stderr, s.Transcript)
	fmt.Fprintf(stdout, "> ")
	input := bufio.NewScanner(s.Stdin)
	for input.Scan() {
		line := input.Text()
		cmd, err := CmdFromString(line)
		if err != nil {
            fmt.Fprintf(stdout, "> ")
			continue
		}
        if s.DryRun {
            fmt.Fprintf(stdout, "%s\n> ", line)
            continue
        }
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintln(stderr, "err", err)
		}
		fmt.Fprintf(stdout, "%s> ", out)
	}
    fmt.Fprintln(stdout, "\nSee you later!")
}

func Main() {
    session := NewSession(os.Stdin, os.Stdout, os.Stderr, )
    transcript, err := os.Create("transcript.txt")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
    defer transcript.Close()
    session.Transcript = transcript
    session.Run()
}