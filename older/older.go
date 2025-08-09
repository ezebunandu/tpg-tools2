package older

import (
	"io/fs"
	"time"
)

func Files(fsys fs.FS, duration time.Duration) (files []string) {
	now := time.Now()
	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		fileInfo, _ := fs.Stat(fsys, p)
		if !fileInfo.IsDir() && now.Sub(fileInfo.ModTime()) > duration {
			files = append(files, p)
		}
		return nil
	})
	return files
}
