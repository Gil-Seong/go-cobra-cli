package cmd

import (
	"fmt"
	cc "go-cli/conf/cobra"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"
var rootCli string = "fc"

var rootCmd = &cobra.Command{
	Version: version,
	Use:     rootCli,
	Short:   rootCli + " app to demonstrate cobra",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {

	cc.Init(&cc.Config{
		RootCmd:       rootCmd,
		Headings:      cc.HiCyan + cc.Bold + cc.Underline,
		Commands:      cc.HiYellow + cc.Bold,
		Aliases:       cc.Bold + cc.Italic,
		CmdShortDescr: cc.HiRed,
		Example:       cc.Italic,
		ExecName:      cc.Bold,
		Flags:         cc.Bold,
		FlagsDescr:    cc.HiRed,
		FlagsDataType: cc.Italic,
	})

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
