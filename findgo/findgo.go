package findgo

import (
    "io/fs"
    "os"
    "path/filepath"
)

func Files(dir string) (paths []string) {
    fsys := os.DirFS(dir)
    fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
        if filepath.Ext(p) == ".go" {
            paths = append(paths, p)
        }
        return nil
    })
    return paths
}