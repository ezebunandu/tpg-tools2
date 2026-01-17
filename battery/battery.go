package battery

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

type Status struct {
	ChargePercent int
}

var pmsetOutput = regexp.MustCompile("([0-9]+)%")

func ParsePmSetOutput(text string) (Status, error) {
	matches := pmsetOutput.FindStringSubmatch(text)
	if len(matches) < 2 {
		return Status{}, fmt.Errorf("failed to parse pmset output: %q", text)
	}
	charge, err := strconv.Atoi(matches[1])
	if err != nil {
		return Status{}, err
	}
	return Status{ChargePercent: charge}, nil
}

func GetPmsetOutput() (string, error){
	data, err := exec.Command("/usr/bin/pmset", "-g",  "ps").CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(data), nil
}