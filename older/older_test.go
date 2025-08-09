package older_test

import (
	"testing"
	"testing/fstest"
	"time"

	"github.com/ezebunandu/older"
    "github.com/google/go-cmp/cmp"
)

func TestFiles__ReturnsFilesOlderThanGivenDuration(t *testing.T){
    t.Parallel()
    now := time.Now()
    fsys := fstest.MapFS{
        "file.go": {ModTime: now},
        "subfolder/subfolder.go": {ModTime: now.Add(-5 * time.Minute)},
        "subfolder2/another.go": {ModTime: now},
        "subfolder2/file.go":{ModTime: now.Add(-5 * time.Minute)},
    }
    want := []string{
        "subfolder/subfolder.go",
        "subfolder2/file.go",
    }
    got := older.Files(fsys, 3 * time.Minute)
    if !cmp.Equal(want, got){
        t.Error(cmp.Diff(want, got))
    }
}