package main

import (
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/a-hilaly/cs/pkg/cmd"
)

var (
	isGit            bool
	useGitIgnoreFile bool
	exclude          []string
)

func init() {
	Cmd.Flags().BoolVarP(&isGit, "is-git", "g", true, "is true will lead program to ignore all files under .git directory")
	Cmd.Flags().BoolVarP(&useGitIgnoreFile, "use-gitignore", "G", true, "i .git directory")
}

func main() {
	if err := Cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

var Cmd = &cobra.Command{
	Use:   "cs",
	Short: "code statistics",
	Run: func(c *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalf("no targets provided")
		}

		for _, target := range args {
			info, err := os.Stat(target)
			if err != nil {
				log.Fatalf("%v\n", err)
			}

			var tw *tablewriter.Table
			if info.IsDir() {
				table, err := cmd.WalkAndParseDir(os.Stdout, target)
				if err != nil {
					log.Fatalf("%v\n", err)
				}

				tw = table
			} else {
				table, err := cmd.ParseFile(os.Stdout, target)
				if err != nil {
					log.Fatalf("%v\n", err)
				}

				tw = table
			}

			tw.Render()

		}
	},
}
