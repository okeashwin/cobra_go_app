package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds any two given integers",
	Long:  "Adds any two integers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add command invoked")
	},
}