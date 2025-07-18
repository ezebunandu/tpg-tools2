package hello_test

import (
	"bytes"
	"testing"

	"github.com/ezebunandu/hello"
)

func TestPrintTo__PrintsHelloMessageToGivenWriter(t *testing.T){
    t.Parallel()
    buf := new(bytes.Buffer)
    hello.PrintTo(buf)
    want := "hello world\n"
    got := buf.String()
    if want != got {
        t.Fatalf("want %q, got %q", want, got) 
    }
}