package count_test

import (
	"bytes"
	"testing"

	"github.com/ezebunandu/count"
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