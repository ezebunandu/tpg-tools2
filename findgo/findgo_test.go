package findgo_test

import (
	"archive/zip"
	"os"
	"testing"
	"testing/fstest"

	"github.com/ezebunandu/findgo"
	"github.com/google/go-cmp/cmp"
)

func TestFiles__CorrectlyListsFilesInTree(t *testing.T){
    t.Parallel()
    fsys := os.DirFS("testdata/tree")
    want := []string{
        "file.go",
        "subfolder/subfolder.go",
        "subfolder2/another.go",
        "subfolder2/file.go",
    }
    got := findgo.Files(fsys)
    if !cmp.Equal(want, got){
        t.Error(cmp.Diff(want, got))
    }
}

func BenchmarkFilesOnDisk(b *testing.B){
    fsys := os.DirFS("testdata/tree")
    for b.Loop(){
        _ = findgo.Files(fsys)
    }
}

func BenchmarkFilesInMemory(b *testing.B){
    fsys := fstest.MapFS{
        "file.go": {},
        "subfolder/subfolder.go": {},
        "subfolder2/another.go": {},
        "subfolder2/file.go": {},
    }
    for b.Loop(){
        _ = findgo.Files(fsys)
    }
}

func TestFiles__CorrectlyListsFilesInMapFS(t *testing.T){
    t.Parallel()
    fsys := fstest.MapFS{
        "file.go": {},
        "subfolder/subfolder.go": {},
        "subfolder2/another.go": {},
        "subfolder2/file.go": {},
    }
    want := []string{
        "file.go",
        "subfolder/subfolder.go",
        "subfolder2/another.go",
        "subfolder2/file.go",
    }
    got := findgo.Files(fsys)
    if !cmp.Equal(want, got){
        t.Error(cmp.Diff(want, got))
    }
}

func TestFiles__CorrectlyListFilesinZIPArchive(t *testing.T){
    t.Parallel()
    fsys, err := zip.OpenReader("testdata/files.zip")
    if err != nil {
        t.Fatal(err)
    }
    want := []string{
        "tree/file.go",
        "tree/subfolder/subfolder.go",
        "tree/subfolder2/another.go",
        "tree/subfolder2/file.go",
    }
    got := findgo.Files(fsys)
    if !cmp.Equal(want, got){
        t.Error(cmp.Diff(want, got))
    }
}