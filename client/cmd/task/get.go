package task

import (
	"github.com/spf13/cobra"
	task_utils "longabonga.com/cobra/template/client/pkg/task"
	"longabonga.com/cobra/template/client/utils"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get a new task",
	Long:  `get a new task`,
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := utils.GetStringFlag(cmd, taskIDFlag)
		if err != nil {
			return err
		}
		return task_utils.GetByID(id)
	},
}

func init() {
	utils.SetRequiredStringFlag(getCmd, taskIDFlag, taskIDShortFlag, taskIDDesc)

	TaskCmd.AddCommand(getCmd)
}
