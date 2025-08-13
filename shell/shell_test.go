package shell_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
    "github.com/ezebunandu/shell"
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