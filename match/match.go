package match

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type matcher struct {
    input io.Reader
    output io.Writer
    text string
}

type option func(* matcher) error

func WithInput(input io.Reader) option {
    return  func(m *matcher) error {
        if input == nil {
            return  fmt.Errorf("nil innput supplied")
        }
        m.input = input
        return  nil
    }
}

func WithOutput(output io.Writer) option {
    return  func(m *matcher) error {
        if output == nil {
            return  fmt.Errorf("nil output supplied")
        }
        m.output = output
        return  nil
    }
}

func WithSearchStringFromArgs(args []string) option {
    return  func(m *matcher) error {
        if len(args) < 1 {
            return fmt.Errorf("emtpy args")
        }
        m.text = args[0]
        return  nil
    }
}

func NewMatcher(opts...option) (*matcher, error){
    m := &matcher{
        input: os.Stdin,
        output: os.Stderr,
    }
    for _, opt := range opts {
        err := opt(m)
        if err != nil {
            return nil, err
        }
    }
    return m, nil
}

func (m *matcher) PrintMatchingLines(){
    input := bufio.NewScanner(m.input)
    for input.Scan(){
        if strings.Contains(input.Text(), m.text){
            fmt.Println(input.Text())
        }
    }
}

func Main(){
    m, err := NewMatcher(
        WithSearchStringFromArgs(os.Args[1:]),
    )
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
    m.PrintMatchingLines()
}

