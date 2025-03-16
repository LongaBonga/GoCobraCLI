package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"longabonga.com/cobra/template/client/cmd/task"
)

var rootCmd = &cobra.Command{
	Use:   "template-cli",
	Short: "A template CLI for",
	Long:  `A template CLI for demo`,
}

func addSubcommandPalletes() {
	rootCmd.AddCommand(task.TaskCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}

func init() {
	addSubcommandPalletes()
}
