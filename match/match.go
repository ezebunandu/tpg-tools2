package match

import (
	"bufio"
	"errors"
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

type option func (*matcher) error 

func NewMatcher(opts...option) (*matcher, error) {
    c := &matcher{
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
    return  func(m *matcher) error {
        if input == nil {
            return  errors.New("nil input reader")
        }
        m.input= input
        return  nil
    }
}

func WithOutput(output io.Writer) option {
    return  func(m *matcher) error {
        if output == nil {
            return  errors.New("nil output writer")
        }
        m.output = output
        return  nil
    }
}

func WithSearchStringFromArgs(args []string) option {
    return func(m *matcher) error {
        if len(args) < 1 {
            return  nil
        }
        m.text = args[0]
        return  nil
    }
}

func (m *matcher) MatchingLines() string {
    var matches string
    input := bufio.NewScanner(m.input)
    for input.Scan() {
        line := string(input.Bytes())
        if strings.Contains(line, m.text) {
            matches = matches + line + "\n"
        }
    }
    return matches
}

func Main(){
    m, err := NewMatcher(
        WithSearchStringFromArgs(os.Args[1:]),
    )
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
    fmt.Println(m.MatchingLines())
}