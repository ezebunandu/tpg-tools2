package count_test

import (
	"bytes"
	"testing"

	"github.com/ezebunandu/count"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestLines__CountsLinesInInput(t *testing.T){
    t.Parallel()

    inputBuf := bytes.NewBufferString("1\n2\n3")
    c, err := count.NewCounter(
        count.WithInput(inputBuf),
    )
    if err != nil {
        t.Fatal(err)
    }
    want := 3
    got := c.Lines()
    if want != got {
        t.Errorf("want %d, got %d", want, got)
    }
}

func TestWithInputFromArgs__SetsInputToGivenPath(t *testing.T){
    t.Parallel()
    args := []string{"testdata/three_lines.txt"}
    c, err := count.NewCounter(
        count.WithInputFromArgs(args),
    )
    if err != nil {
        t.Fatal(err)
    }
    want := 3
    got := c.Lines()
    if want != got {
        t.Errorf("want %d, got %d", want, got)
    }
}

func TestWithInputFromArgs__ReadsInputFromMultipleArgs(t *testing.T){
    t.Parallel()
    args := []string{"testdata/three_lines.txt", "testdata/two_lines.txt"}
    c, err := count.NewCounter(
        count.WithInputFromArgs(args),
    )
    if err != nil {
        t.Fatal(err)
    }
    want := 5
    got := c.Lines()
    if want != got {
        t.Errorf("want %d, got %d", want, got)
    }
}

func TestWithInputFromArgs__IgnoresEmptyArgs(t *testing.T){
    t.Parallel()
    inputBuf := bytes.NewBufferString("1\n2\n3")
    c, err := count.NewCounter(
        count.WithInput(inputBuf),
        count.WithInputFromArgs([]string{}),
    )
    if err != nil {
        t.Fatal(err)
    }
    want := 3 
    got := c.Lines()
    if want != got {
        t.Errorf("want %d, got %d", want, got)
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
        "count": count.Main,
    })
}