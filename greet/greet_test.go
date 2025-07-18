package greet_test

import (
	"bytes"
	"testing"

	"github.com/ezebunandu/greet"
)

func TestGreetUser__PromptsUserForNameAndRendersGreeting(t *testing.T){
    t.Parallel()
    input := bytes.NewBufferString("Sam")
    output := new(bytes.Buffer)
    greet.GreetUser(input, output)
    want := "What is your name?\nHello, Sam"
    got := output.String()
    if want != got {
        t.Fatalf("want %q, got %q", want, got) 
    }
}