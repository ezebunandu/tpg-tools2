package greet

import (
	"bufio"
	"fmt"
	"io"
	"os"
)


func GreetUser(r io.Reader, w io.Writer){
	name := "you"
	fmt.Fprintln(w, "What is your name?")
	input := bufio.NewScanner(r)
	if input.Scan() {
		name = input.Text()
	}
	fmt.Fprint(w, "Hello, ", name)
}

func Main(){
	GreetUser(os.Stdin, os.Stdout)
}