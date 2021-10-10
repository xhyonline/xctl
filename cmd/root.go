package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// 跟命令
var rootCmd = &cobra.Command{
	Short: "help",
	Long:  `这是一份帮助手册,您可以查看命令`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init()  {
	rootCmd.AddCommand(createCmd)
}

func Execute() {
	initArgs()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("出错了", err)
		os.Exit(1)
	}
}