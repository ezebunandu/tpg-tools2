package older

import (
	"io/fs"
	"time"
)

func Files(fsys fs.FS, age time.Duration) (paths []string) {
    threshold := time.Now().Add(-age)
	fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		info, err := d.Info()
		if err != nil {
			return err
		}
		if info.ModTime().Before(threshold){
			paths = append(paths, path)
		}
		return nil
	})
	return paths
}
