package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xhyonline/xctl/internal"
)

var protoCmd = &cobra.Command{
	Use:     "gen-proto",
	Short:   "根据 protobuf 文件生成 pb 文件",
	Long:    `编译 protobuf `,
	Example: "xctl gen-proto",
	Run: func(cmd *cobra.Command, _ []string) {
		if err := internal.GeneratePb(); err != nil {
			fmt.Println(err)
		}
	},
}
