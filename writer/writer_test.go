package writer_test

import (
	"os"
	"testing"

	"github.com/ezebunandu/writer"
	"github.com/google/go-cmp/cmp"
    "github.com/rogpeppe/go-internal/testscript"
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

func TestWriteToFile__ReturnsErrorForUnwritableFile(t *testing.T){
    t.Parallel()
    path := "bogus/write_test.txt"
    err := writer.WriteToFile(path, []byte{})
    if err == nil {
        t.Fatal("want error when file not writable, got none")
    }
}

func TestWriteToFile__ClobbersExistingFile(t *testing.T){
    t.Parallel()
    path := t.TempDir() + "/write_test.txt"
    err := os.WriteFile(path, []byte{4, 5, 6}, 0o600)
    if err != nil {
        t.Fatal(err)
    }
    want := []byte{1, 2, 3}
    err = writer.WriteToFile(path, want)
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

func TestWriteToFile__ChangesPermsOnExistingFile(t *testing.T){
    t.Parallel()
    path := t.TempDir() + "/perm_test.txt"
    err := os.WriteFile(path, []byte{}, 0o644)
    if err != nil {
        t.Fatal(err)
    }
    err = writer.WriteToFile(path, []byte{1, 2, 3})
    if err != nil {
        t.Fatal(err)
    }
    stat, err := os.Stat(path)
    if err != nil {
        t.Fatal(err)
    }
    want := os.FileMode(0o600)
    got := stat.Mode().Perm()
    if want != got {
        t.Errorf("want file mode %q, got %q", want, got)
    }
}

func Test(t *testing.T){
    t.Parallel()
    testscript.Run(t, testscript.Params{
        Dir: "testdata/script",
    })
}

func TestMain(m *testing.M) {
    testscript.Main(m, map[string]func(){
        "writefile": writer.Main,
    })
}