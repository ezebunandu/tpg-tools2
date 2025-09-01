package pipeline

import (
	"io"
	"os"
	"strings"
)

type Pipeline struct{
    Reader io.Reader
    Output io.Writer
    Error error
}

func New() *Pipeline {
    return  &Pipeline{
        Output: os.Stdout,
    }
}

func FromString(s string) *Pipeline{
    p := New()
    p.Reader = strings.NewReader(s)
    return p
}

func FromFile(pathname string) *Pipeline{
    p := New()
    f, err := os.Open(pathname)
    if err != nil {
        p.Error = err
        return  p
    }
    p.Reader = f
    return  p
}

func (p *Pipeline) Stdout(){
    if p.Error != nil {
        return
    }
    io.Copy(p.Output, p.Reader)
}

func (p *Pipeline) String() (string, error){
    if p.Error != nil {
        return "", p.Error
    }
    data, err := io.ReadAll(p.Reader)
    if err!= nil {
        return  "", err
    }
    return string(data), nil
}
