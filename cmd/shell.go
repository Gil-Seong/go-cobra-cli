package cmd

import (
	"fmt"
	"os/exec"
	"os"
	"github.com/spf13/cobra"
)

var sh_command string

// getCmd represents the get command
var shellCmd = &cobra.Command{
	Use:   "shell",
	Aliases: []string{"s"},
	// Args: cobra.ExactArgs(1), //args 로 받을 때 필수 값 사용
	Short: "shell brief description of your command",
	Example: rootCli+` shell --command "[command]"
`+rootCli+` shell -c "[command]"`,
	Run: func(cmd *cobra.Command, args []string) {

		// kernelInfo, err := exec.Command("uname", "-r").Output()
		// check(err)
		// fmt.Print("Kernel Version : ", string(kernelInfo))

		com := exec.Command(sh_command)
		com.Stdout = os.Stdout

		if err := com.Run(); err != nil {
			fmt.Println(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(shellCmd)

	shellCmd.PersistentFlags().StringVarP(&sh_command, "command", "c", "", "shellCmd help for shell")
  shellCmd.MarkFlagRequired("command") //필수값 지정

}
