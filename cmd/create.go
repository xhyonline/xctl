package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xhyonline/xctl/internal"
)

var args = new(internal.CreateArgs)

func initArgs() {
	createCmd.Flags().StringVar(&args.AppName, "with-name", "", "必填项 应用名称,例如:myapp,")
	createCmd.Flags().StringVar(&args.Mod, "with-mod", "", "必填项 初始化 go mod 例如: github.com/myapp")
	createCmd.Flags().BoolVar(&args.WithHTTPServer, "with-http-server", false,
		"二选一或全选(全选则附带grpc-server附带GRPC Client 说明文档 rpc 目录下) 必填项 是否是一个 HTTP 服务?")
	createCmd.Flags().BoolVar(&args.WithGRPCServer, "with-grpc-server", false, "二选一 必填项 是否是一个 GRPC 服务?")
	createCmd.Flags().BoolVar(&args.WithMySQL, "with-mysql", false, "是否使用 mysql 数据库")
	createCmd.Flags().BoolVar(&args.WithRedis, "with-redis", false, "是否使用 redis 缓存")
	createCmd.Flags().BoolVar(&args.WithGithubAction, "with-githubAction", false, "初始化 github action 集成。注:生成的 " +
		"action.yml 文件中采用两种发布模式,你需要自行选取一种。第一种是在 github 上构建,并使用 rsync 进行发布打包后的二进制执行文件。第二种是使用 rsync 同步代码到自己的云服务器,并进行构建。" +
		"由于 Github 自身网速很慢,你需要考量自己的项目适合哪一种模式")
}

var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "创建项目",
	Long:    `你可以根据自己的需求创建一个项目,示例如下`,
	Example: "xctl create --with-name myapp --with-http-server --with-mod github.com/xhyonline/myapp ",
	Run: func(cmd *cobra.Command, _ []string) {
		if args.AppName == "" || args.Mod == "" || (!args.WithHTTPServer && !args.WithGRPCServer) {
			fmt.Println("必填参数不能为空:" + "  --with-name" + "  --with-mod" + "--with-grpc-server" + "--with-grpc-server")
			_ = cmd.Help()
			return
		}
		internal.CreateProject(args)
	},
}
