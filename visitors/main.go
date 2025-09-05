package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/bitfield/script"
)

func main(){
    re := regexp.MustCompile(".go")
    n, err := script.FindFiles(".").MatchRegexp(re).Concat().CountLines()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(n)
}