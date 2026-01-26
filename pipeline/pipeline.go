package pipeline

import (
    "io"
	"strings"
)

type Pipeline struct {
	Reader io.Reader
	Output io.Writer
	Error  error
}

func FromString(s string) *Pipeline {
	return &Pipeline{Reader: strings.NewReader(s)}
}

func (p Pipeline) Stdout() {
	if p.Error != nil {
		return
	}
	io.Copy(p.Output, p.Reader)
}
