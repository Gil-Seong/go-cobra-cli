package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var str string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get str : ", str)
    fmt.Println("get args : ", args)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.PersistentFlags().StringVar(&str, "str", "", "ACmd help for boo")
  getCmd.MarkFlagRequired("str") //필수값 지정

  getCmd.Flags().BoolP("list", "l", false, "get list")
}
