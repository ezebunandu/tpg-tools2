package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ezebunandu/shell"
)

func main(){
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        line := input.Text()
        cmd, err := shell.CmdFromString(line)
        if err != nil {
            continue
        }
        out, err := cmd.CombinedOutput()
        if err != nil {
            fmt.Println("err", err)
        }
        fmt.Printf("%s", out)
    }
}