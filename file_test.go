package pkgmv

import (
	"fmt"
	"testing"
)

func TestGetFiles(t *testing.T) {
	fis, _ := GetGoFiles(srcpath)
	fmt.Println(len(fis))
}
