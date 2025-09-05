package main

import (

	"github.com/bitfield/script"
)

func main(){
    script.File("log.txt").Column(1).Freq().First(10).Stdout()
}