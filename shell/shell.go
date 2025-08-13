package shell

import (
	"fmt"
	"os/exec"
	"strings"
)

// "os/exec"

type cmd struct {
    Args []string
}
func CmdFromString(input string) (*exec.Cmd, error) {
    args := strings.Fields(input)
    if len(args) < 1 {
        return  nil, fmt.Errorf("empty input %s", input )
    }
    return exec.Command(args[0], args[1:]...), nil
}