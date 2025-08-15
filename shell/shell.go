package shell

import (
    "bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)


type cmd struct {
    Args []string
}

type Session struct {
    Stdin io.Reader
    Stdout io.Writer
    Stderr io.Writer
    DryRun bool
    Transcript io.Writer
}

func NewSession(input io.Reader, output, err io.Writer) *Session {
    return &Session{Stdin: input, Stdout: output, Stderr: err, Transcript: io.Discard}
}

func CmdFromString(input string) (*exec.Cmd, error) {
    args := strings.Fields(input)
    if len(args) < 1 {
        return  nil, fmt.Errorf("empty input %s", input )
    }
    return exec.Command(args[0], args[1:]...), nil
}

func (session *Session) Run() {
    stdout := io.MultiWriter(session.Stdout, session.Transcript)
    stderr := io.MultiWriter(session.Stderr, session.Transcript)
    fmt.Fprintf(session.Stdout, "> ")
    input := bufio.NewScanner(session.Stdin)
    for input.Scan(){
        line := input.Text()
        cmd, err := CmdFromString(line)
        if err != nil {
            fmt.Fprintf(stdout, "> ")
            continue
        }
        if session.DryRun {
            fmt.Fprintf(stdout, "%s\n> ", line)
            continue
        }
        output, err := cmd.CombinedOutput()
        if err != nil {
            fmt.Println(stderr, "error:", err)
        } 
        fmt.Fprintf(stdout, "%s>", output)
    }
    fmt.Fprintln(stdout, "\nBe seeing you!")
}

func Main() {
    session := NewSession(os.Stdin, os.Stdout, os.Stderr)
    transcript, err := os.Create("transcript.txt")
    if err != nil {
        fmt.Fprint(os.Stderr, err)
        os.Exit(1)
    }
    defer transcript.Close()
    session.Transcript = transcript
    session.Run()
}