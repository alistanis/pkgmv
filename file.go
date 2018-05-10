package pkgmv

import (
	"os"
	"sync"

	"path/filepath"

	"github.com/MichaelTJones/walk"
)

type FileInfo struct {
	Path     string
	FileInfo os.FileInfo
	Err      error
}

var srcpath = filepath.Join(os.Getenv("GOPATH"), "src")

func GetGoFiles(path string) ([]FileInfo, error) {

	if _, err := os.Stat(path); err != nil {
		return nil, err
	}

	var mu sync.Mutex

	var fis []FileInfo

	walk.Walk(path, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) != ".go" {
			return nil
		}
		mu.Lock()
		fis = append(fis, FileInfo{path, info, err})
		mu.Unlock()
		return nil
	})

	return fis, nil
}
