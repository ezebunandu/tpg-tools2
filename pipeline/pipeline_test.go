package pipeline_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/ezebunandu/pipeline"
	"github.com/google/go-cmp/cmp"
)

func TestStdout_PrintsMessageToOutput(t *testing.T) {
    t.Parallel()
    want := "Hello, world\n"
    p := pipeline.FromString(want)
    buf := new(bytes.Buffer)
    p.Output = buf
    p.Stdout()
    if p.Error != nil {
        t.Fatal(p.Error)
    }
    got := buf.String()
    if !cmp.Equal(want, got) {
        t.Errorf("want %q, got %q", want, got)
    }
}

func TestStdout_PrintsNothingOnError(t *testing.T) {
    t.Parallel()
    p := pipeline.FromString("Hello, world\n")
    p.Error = errors.New("oh no")
    buf := new(bytes.Buffer)
    p.Output = buf
    p.Stdout()
    got := buf.String()
    if got != "" {
        t.Errorf("want no output from Stdout after error, but got %q", got)
    }
    
}