package count

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type counter struct {
    input io.Reader 
    output io.Writer
    files []io.Reader
}

type option func (*counter) error 

func NewCounter(opts...option) (*counter, error) {
    c := &counter{
        input: os.Stdin,
        output: os.Stdout,
    }
    for _, opt := range opts {
        err := opt(c)
        if err != nil {
            return nil, err
        }
    }
    return  c, nil
}

func WithInput(input io.Reader) option {
    return  func(c *counter) error {
        if input == nil {
            return  errors.New("nil input reader")
        }
        c.input= input
        return  nil
    }
}

func WithOutput(output io.Writer) option {
    return  func(c *counter) error {
        if output == nil {
            return  errors.New("nil output writer")
        }
        c.output = output
        return  nil
    }
}

func WithInputFromArgs(args []string) option {
    return func(c *counter) error {
        if len(args) < 1 {
            return  nil
        }
        c.files = make([]io.Reader, len(args))
        for i, path := range args {
            f, err := os.Open(path)
            if err != nil {
                return  err
            }
            c.files[i] = f
        }
        c.input = io.MultiReader(c.files...)
        return  nil
    }
}

func (c *counter) Lines() int {
    var lines int
    input := bufio.NewScanner(c.input)
    for input.Scan() {
        lines++
    }
    for _, f := range c.files {
        f.(io.Closer).Close()
    }
    return  lines
}

func Main() {
    c, err := NewCounter(
        WithInputFromArgs(os.Args[1:]),
    )
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
    fmt.Println(c.Lines())
}