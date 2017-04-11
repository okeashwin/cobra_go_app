package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(subtractCmd)
}

var subtractCmd = &cobra.Command{
	Use:   "subtract",
	Short: "Subtracts any two given integers",
	Long:  "Subtracts any two integers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Subtract command invoked")
		if len(args) != 2 {
			fmt.Println("Insufficient number of arguments provided")
			os.Exit(-1)
		}

		fmt.Println("args length :", len(args))
		num1, err1 := strconv.Atoi(args[0])
		if err1 != nil {
			fmt.Printf("Couldn't convert %s to a number\n", args[0])
			os.Exit(-1)
		}
		num2, err2 := strconv.Atoi(args[1])
		if err2 != nil {
			fmt.Printf("Couldn't convert %s to a number\n", args[1])
			os.Exit(-1)			
		}
		fmt.Println("Result is ", num1-num2)
		os.Exit(0)
	},
}