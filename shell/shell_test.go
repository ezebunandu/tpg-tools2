package shell_test

import (
	"os"
	"testing"

	"github.com/ezebunandu/shell"
	"github.com/google/go-cmp/cmp"
)

func TestCmdFromString__CreatesExpectedCmd(t *testing.T){
    t.Parallel()
    input := "/bin/ls -l main.go"
    cmd, err := shell.CmdFromString(input)
    if err != nil {
        t.Fatal(err)
    }
    want := []string{"/bin/ls", "-l", "main.go"}
    got := cmd.Args
    if !cmp.Equal(want, got){
        t.Error(cmp.Diff(want, got))
    }
}

func TestCmdFromString__ErrorsOnEmptyInput(t *testing.T){
    t.Parallel()
    input := ""
    _, err := shell.CmdFromString(input)
    if err == nil {
        t.Fatal("want error on empty input, got nil")
    }
}

func TestNewSession__CreatesExpectedSession(t *testing.T){
    t.Parallel()
    want := shell.Session{
        Stdin: os.Stdin,
        Stdout: os.Stdout,
        Stderr: os.Stderr,
    }
    got := *shell.NewSession(os.Stdin, os.Stdout, os.Stderr)
    if want != got {
        t.Errorf("want %#v, got %#v", want, got)
    }
}