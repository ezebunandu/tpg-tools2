package pipeline_test

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/ezebunandu/pipeline"
	"github.com/google/go-cmp/cmp"
)

func TestStdout__PrintsMessageToOutput(t *testing.T) {
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

func TestStdout__PrintsNothingOnError(t *testing.T) {
    t.Parallel()
    p := pipeline.FromString("Hello, world\n")
    p.Error = errors.New("Oops")
    buf := new(bytes.Buffer)
    p.Output = buf
    p.Stdout()
    got := buf.String()
    if got != ""{
        t.Errorf("want no output from Stdout after error, got %q", got)
    }
}

func TestFromFile__ReadsAllDataFromFile(t *testing.T) {
    t.Parallel()
    want := []byte("hello, world\n")
    p := pipeline.FromFile("testdata/hello.txt")
    if p.Error != nil {
        t.Fatal(p.Error)
    }
    got, err := io.ReadAll(p.Reader)
    if err != nil {
        t.Fatal(err)
    }
    if !cmp.Equal(want, got) {
        t.Errorf("want %q, got %q", want, got)
    }
}

func TestFromFile__SetsErrorGivenNonexistentFile(t *testing.T){
    t.Parallel()
    p := pipeline.FromFile("does-not-exist.txt")
    if p.Error == nil {
        t.Fatal("want error opening non existent file but got nil")
    }
}

func TestString__ReturnsPipelineContents(t *testing.T) {
    t.Parallel()
    want := "Hello, world\n"
    p := pipeline.FromString(want)
    got, err := p.String()
    if err != nil {
        t.Fatal(err)
    }
    if !cmp.Equal(want, got) {
        t.Errorf("want %q, got %q", want, got)
    }
}

func TestString__ReturnsErrorWhenPipeErrorSet(t *testing.T){
    t.Parallel()
    p := pipeline.FromString("Hello, world\n")
    p.Error = errors.New("Oops")
    _, err := p.String()
    if err == nil {
        t.Error("want error when pipe error set but got none")
    } 
}