package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xhyonline/xctl/internal"
)

var args = new(internal.Args)

func initArgs() {
	createCmd.Flags().StringVar(&args.AppName, "with-appName", "myapp", "应用名称,默认为 myapp")
	createCmd.Flags().StringVar(&args.Mod, "with-mod", "github.com/myapp", "初始化 go mod 默认为 github.com/myapp")
	createCmd.Flags().BoolVar(&args.WithHTTPServer, "with-http-server", true, "是否是一个 HTTP 服务? 默认是")
	createCmd.Flags().BoolVar(&args.WithMySQL, "with-mysql", false, "是否使用 mysql 数据库")
	createCmd.Flags().BoolVar(&args.WithRedis, "with-redis", false, "是否使用 redis 缓存")
	createCmd.Flags().BoolVar(&args.WithEtcd, "with-etcd", false, "是否使用etcd")
	createCmd.Flags().BoolVar(&args.WithGithubAction, "with-githubAction", false, "是否初始化 github action 集成")
}

var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "创建项目",
	Long:    `你可以根据自己的需求创建一个项目,示例如下`,
	Example: "xhyctl build --with-db --with-redis",
	Run: func(cmd *cobra.Command, _ []string) {
		internal.CreateProject(args)
	},
}
