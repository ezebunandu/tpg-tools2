package main

import "github.com/bitfield/script"

func main(){
    script.Args().Concat().Column(1).Freq().First(10).Stdout()
}