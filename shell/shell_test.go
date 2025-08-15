package shell_test

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/ezebunandu/shell"
	"github.com/google/go-cmp/cmp"
    "github.com/rogpeppe/go-internal/testscript"
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
        Transcript: io.Discard,
    }
    got := *shell.NewSession(os.Stdin, os.Stdout, os.Stderr)
    if want != got {
        t.Errorf("want %#v, got %#v", want, got)
    }
}

func TestRun__ProducesExpectedOutput(t *testing.T){
    t.Parallel()
    in := strings.NewReader("echo hello\n\n")
    out := new(bytes.Buffer)
    session := shell.NewSession(in, out, io.Discard)
    session.DryRun = true
    session.Run()
    want := "> echo hello\n> > \nBe seeing you!\n"
    got := out.String()
    if !cmp.Equal(want, got){
        t.Error(cmp.Diff(want, got))
    }
}


func Test(t *testing.T){
    t.Parallel()
    testscript.Run(t, testscript.Params{
        Dir: "testdata/script",
    })
}

func TestMain(m *testing.M) {
    testscript.Main(m, map[string]func(){
        "shell": shell.Main,
    })
}