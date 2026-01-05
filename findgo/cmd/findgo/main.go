package main

import (
	"fmt"
	"os"

    "github.com/ezebunandu/findgo"
)

func main(){
    Usage := `findgo <path>`
    if len(os.Args) < 2 {
        fmt.Println(Usage)
        os.Exit(0)
    }
    fsys := os.DirFS(os.Args[1])
    paths := findgo.Files(fsys)
    for _, p := range paths {
        fmt.Println(p)
    }
}