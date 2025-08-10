package older

import (
	"io/fs"
	"time"
)

func Files(fsys fs.FS, age time.Duration) (paths []string) {
	threshold := time.Now().Add(-age)
	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		fi, err  := d.Info()
		if err != nil || fi.IsDir() {
			return nil
		}
		if fi.ModTime().Before(threshold){
			paths = append(paths, p)
		}
		return nil
	})
	return paths
}
