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
	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestFromFile_SetsErrorGivenNonExistentFile(t *testing.T) {
	t.Parallel()
	p := pipeline.FromFile("nonexistent.txt")
	if p.Error == nil {
		t.Fatal("want error openining non-existent file, got nil")
	}
}

func TestString_ReturnsPipeContents(t *testing.T) {
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

func TestString_ReturnsErrorWhenPipeErrorSet(t *testing.T) {
	t.Parallel()
	p := pipeline.FromString("Hello, world\n")
	p.Error = errors.New("oh no")
	_, err := p.String()
	if err == nil {
		t.Error("want error from String when pipeline has error, but got nil")
	}
}

func TestColumn_SelectsExpectedColumn(t *testing.T) {
	t.Parallel()
	input := "1 2 3\n1 2 3\n1 2 3\n"
	p := pipeline.FromString(input)
	want := "2\n2\n2\n"
	got, err := p.Column(2).String()
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestColumn_ProducesNothingWhenPipeErrorSet(t *testing.T) {
	t.Parallel()
	p := pipeline.FromString("1 2 3\n")
	p.Error = errors.New("oh no")
	data, err := io.ReadAll(p.Column(1).Reader)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) > 0 {
		t.Errorf("want no output from Column after error, but got %q", data)
	}
}

func TestColumn_SetsErrorAndProducesNothingGivenInvalidArg(t *testing.T) {
	t.Parallel()
	p := pipeline.FromFile("1 2 3\n1 2 3\n1 2 3\n")
	p.Column(-1)
	if p.Error == nil {
		t.Fatalf("want error given invalid column but got none")
	}
	data, err := io.ReadAll(p.Column(1).Reader)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) > 0 {
		t.Errorf("want no output from Column with invalid col, but got %q", data)
	}
}
