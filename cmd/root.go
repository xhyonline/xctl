package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const version = "v0.1.2021111715"

// rootCmd 根命令
var rootCmd = &cobra.Command{
	Short: "help",
	Long:  `这是一份帮助手册,您可以查看命令`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

// versionCmd 显示当前版本
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version",
	Long:  `example: xctl version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("当前版本号:" + version)
	},
}

func init() {
	rootCmd.AddCommand(createCmd, versionCmd, protoCmd)
}

func Execute() {
	initArgs()
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("出错了 %s\n", err)
		os.Exit(1)
	}
}
