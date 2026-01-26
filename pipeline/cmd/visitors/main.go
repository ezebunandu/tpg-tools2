package main

import "github.com/ezebunandu/pipeline"

func main(){
    pipeline.FromString("hello world\n").Stdout()
}