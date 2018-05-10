package pkgmv

import (
	"fmt"
	"os"
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

	if _, err := os.Stat(toPkg); err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(toPkg, 0660)
			if err != nil {
				return err
			} else {
				fmt.Println("made", toPkg)
			}
		}
	}

	// remove this eventually
	err = os.RemoveAll(toPkg)
	if err != nil {
		return err
	}

	fmt.Println(pkgFiles)
	fmt.Println(fromPkg, toPkg)
	return nil
}
