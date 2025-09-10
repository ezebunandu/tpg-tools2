package battery_test

import (
	"os"
	"testing"
    "github.com/ezebunandu/battery"
    "github.com/google/go-cmp/cmp"
)

func TestParsePmsetOutput__GetsChargePercent(t *testing.T){
    t.Parallel()
    data, err := os.ReadFile("testdata/pmset.txt")
    if err != nil {
        t.Fatal(err)
    }
    want := battery.Status{
        ChargePercent: 100,
    }
    got, err := battery.ParsePmsetOutput(string(data))
    if err != nil {
        t.Fatal(err)
    }
    if !cmp.Equal(want, got){
        t.Error(cmp.Diff(want, got))
    }
}

func TestToJSON__GivesExpectedJSON(t *testing.T){
    t.Parallel()
    batt := battery.Battery{
        Name: "InternalBattery-0",
        ID: 10813539,
        ChargePercent: 100,
        TimeToFullCharge: "0.00",
        Present: true,
    }
    wantBytes, err := os.ReadFile("testdata/battery.json")
    if err != nil {
        t.Fatal(err)
    }
    want := string(wantBytes)
    got := batt.ToJSON() + "\n"
    if !cmp.Equal(want, got){
        t.Error(cmp.Diff(want, got))
    }
}