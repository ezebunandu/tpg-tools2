package pipeline_test

import (
	"bytes"
	"errors"
	"io"
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

func TestFromFile_ReadsAllDataFromFile(t *testing.T) {
    t.Parallel()
    want := []byte("Hello, world\n")
    p := pipeline.FromFile("testdata/hello.txt")
    if p.Error != nil {
        t.Fatal(p.Error)
    }
    got, err := io.ReadAll(p.Reader)
    if err != nil {
        t.Fatal(err)
    }
    if !cmp.Equal(want, got){
        t.Errorf("want %q, got %q", want, got)
    }
}

func TestFromFile_SetsErrorGivenNonExistentFile(t *testing.T){
    t.Parallel()
    p := pipeline.FromFile("nonexistent.txt")
    if p.Error == nil {
        t.Fatal("want error openining non-existent file, got nil")
    }
}