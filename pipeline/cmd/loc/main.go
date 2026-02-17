package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/bitfield/script"
)

func main() {
	goFiles := regexp.MustCompile(".go$")
	lines, err := script.FindFiles(".").MatchRegexp(goFiles).Concat().RejectRegexp(regexp.MustCompile(`^\s*$`)).CountLines()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("You have written %d lines of Go in this Project. Nice work!\n", lines)
}
