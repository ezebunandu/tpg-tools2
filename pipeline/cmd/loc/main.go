package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/bitfield/script"
)

func main(){
    re := regexp.MustCompile(".go")
    lines, err := script.FindFiles(".").MatchRegexp(re).Concat().RejectRegexp(regexp.MustCompile(`^\s*$`)).CountLines()
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
    fmt.Printf("You've written %d lines of Go in this project\n", lines)
}