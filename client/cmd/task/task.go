package task

import (
	"github.com/spf13/cobra"
)

const (
	taskIDFlag      = "id"
	taskIDShortFlag = "i"
	taskIDDesc      = "task id"
)

var TaskCmd = &cobra.Command{
	Use:   "task",
	Short: "command to manage tasks",
	Long: `command fot managing tasks. 
	It allows users to create, task`,
}
