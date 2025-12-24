package match_test

import (
	"testing"

	"github.com/ezebunandu/match"
	"github.com/rogpeppe/go-internal/testscript"
)

func Test(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}

func TestMain(m *testing.M) {
	testscript.Main(m, map[string]func(){
		"match": match.Main,
	})
}
