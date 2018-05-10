package pkgmv

import (
	"errors"

	"fmt"

	"path/filepath"

	"github.com/spf13/cobra"
)

// ExecuteRoot executes the root command
func ExecuteRoot() error {
	return setupCommands().Execute()
}

func rootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "pkgmv",
		Short: "move go packages without killing yourself",
	}
}

func errPackageNotFound(err error) error {
	return fmt.Errorf("pkgmv: package to move not found, %s", err)
}

func errPackageAlreadyExists(err error) error {
	return fmt.Errorf("pkgmv: package destination already exists, if you wish to overwrite, remove it, %s", err)
}

func addMoveCmd(root *cobra.Command) {
	var mvOpts moveOptions

	cmd := &cobra.Command{
		Use:   "move [flags] [package] [destination]",
		Short: "move <package> to <new-package>",
		Long:  "moves <package> to <new-package> - paths must be go package names",
		RunE: func(c *cobra.Command, args []string) error {
			if len(args) != 2 {
				return errors.New("must give exactly 2 arguments to move, the package to move and the destination")
			}

			pkg := filepath.Join(srcpath, args[0])
			dest := filepath.Join(srcpath, args[1])

			return move(pkg, dest, mvOpts)
		},
	}

	cmd.Flags().BoolVarP(&mvOpts.DryRun, "dry-run", "d", false, "performs a dry run of the refactor, giving a report of what will be moved")
	cmd.Flags().BoolVarP(&mvOpts.Verbose, "verbose", "v", false, "prints verbose output on a run or dry-run of pkgmv")
	cmd.Flags().BoolVarP(&mvOpts.Copy, "copy", "c", false, "copy files rather than just moving them (import paths in other files will still be changed)")

	root.AddCommand(cmd)
}

func setupCommands() *cobra.Command {
	root := rootCmd()
	addMoveCmd(root)

	return root
}
