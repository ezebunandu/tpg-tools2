package battery

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

type Status struct {
    ChargePercent int
}

type Battery struct {
    Name string
    ID int
    ChargePercent int
    TimeToFullCharge string
    Present bool
}

func (b *Battery) ToJSON() string {
    output, err := json.Marshal(b)
    if err != nil {
        panic(err)
    }
    return string(output)
}

var pmsetOutput = regexp.MustCompile("([0-9]+)%")

func ParsePmsetOutput(text string) (Status, error) {
    matches := pmsetOutput.FindStringSubmatch(text)
    if len(matches) < 2 {
        return Status{}, fmt.Errorf("failed to parse pmset output: %q", text)
    }
    charge, err := strconv.Atoi(matches[1])
    if err != nil {
        return  Status{}, fmt.Errorf("failed to parse charge percentage: %q", matches[1])
    }
    return Status{
        ChargePercent: charge,
    }, nil
}

func GetPmsetOuput() (text string, err error){
    data, err := exec.Command("/usr/bin/pmset", "-g", "ps").CombinedOutput()
    if err != nil {
        return "", nil
    }
    return string(data), nil
}