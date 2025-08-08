package main

import (
	"fmt"
	"os"

	"github.com/ezebunandu/findgo"
)

func main(){
    fsys := os.DirFS(os.Args[1])
    paths := findgo.Files(fsys)
    for _, p := range paths {
        fmt.Println(p)
    }
}