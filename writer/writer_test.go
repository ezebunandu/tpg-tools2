package writer_test

import (
	"os"
	"testing"

	"github.com/ezebunandu/writer"
	"github.com/google/go-cmp/cmp"
)

func TestWriteToFile__WritesGivenDataToFile(t *testing.T){
    t.Parallel()
    path := t.TempDir() + "/write_test.txt"
    want := []byte{1, 2, 3}
    err := writer.WriteToFile(path, want)
    if err != nil {
        t.Fatal(err)
    }
    got, err := os.ReadFile(path)
    if err != nil {
        t.Fatal(err)
    }
    if !cmp.Equal(want, got){
        t.Fatal(cmp.Diff(want, got))
    }
}
