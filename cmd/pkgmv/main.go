package main

import (
	"fmt"
	"os"

	"github.com/alistanis/pkgmv"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	return pkgmv.ExecuteRoot()
}
