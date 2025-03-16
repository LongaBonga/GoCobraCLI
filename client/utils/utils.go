package utils

import (
	"fmt"

	"github.com/spf13/cobra"
)

func SetRequiredStringFlag(cmd *cobra.Command, name, shorthand, description string) {
	cmd.Flags().StringP(name, shorthand, "", description)
	cmd.MarkFlagRequired(name)
}

func GetStringFlag(cmd *cobra.Command, name string) (string, error) {
	value, err := cmd.Flags().GetString(name)
	if err != nil {
		return "", err
	}
	if value == "" {
		return "", fmt.Errorf("%s flag is required", name)
	}
	return value, nil
}

func GetUintFlag(cmd *cobra.Command, name string) (uint, error) {
	return cmd.Flags().GetUint(name)
}

func SetIntFlag(cmd *cobra.Command, name, description string) {
	cmd.Flags().IntP(name, "", 0, description)
}

func GetIntFlag(cmd *cobra.Command, name string) (int, error) {
	return cmd.Flags().GetInt(name)
}
