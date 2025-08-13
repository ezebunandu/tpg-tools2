package shell

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// "os/exec"

type cmd struct {
    Args []string
}

type Session struct {
    Stdin io.Reader
    Stdout io.Writer
    Stderr io.Writer
}

func NewSession(input io.Reader, output, err io.Writer) *Session {
    return &Session{Stdin: input, Stdout: output, Stderr: err}
}

func CmdFromString(input string) (*exec.Cmd, error) {
    args := strings.Fields(input)
    if len(args) < 1 {
        return  nil, fmt.Errorf("empty input %s", input )
    }
    return exec.Command(args[0], args[1:]...), nil
}

func (session *Session) Run() {}

func Main() int {
    session := NewSession(os.Stdin, os.Stdout, os.Stderr)
    session.Run()
    return  0
}