package pkgmv

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type moveOptions struct {
	Verbose bool
	DryRun  bool
	Copy    bool
}

func move(fromPkg, toPkg string, opts moveOptions) error {
	pkgFiles, err := GetGoFiles(fromPkg)
	if err != nil {
		return errPackageNotFound(err)
	}

	if !opts.DryRun {
		if _, err := os.Stat(toPkg); err == nil {
			// the package should never exist
			// we'll simply bail and not worry about overwriting and bad behavior
			return errPackageAlreadyExists
		}
	}

	for _, f := range pkgFiles {
		destinationPath := filepath.Join(toPkg, strings.Replace(f.Path, fromPkg, "", -1))
		if f.FileInfo.IsDir() {
			err = os.MkdirAll(destinationPath, 0660)
			if err != nil {
				return err
			}
			fmt.Println("made", destinationPath)
		}
	}

	// remove this eventually
	err = os.RemoveAll(toPkg)
	if err != nil {
		return err
	}

	return nil
}
