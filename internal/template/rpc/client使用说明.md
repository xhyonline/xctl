```
package rpc

import (
	"github.com/xhyonline/http-framework/component"
	"%goMod/gen/golang"
	"github.com/xhyonline/xutil/grpc"
)

var MyApp golang.RunnerClient

func GetMyApp() golang.RunnerClient {
	if MyApp == nil {
		conn := grpc.NewGRPCClient("myapp", component.Instance.ETCD)
		MyApp = golang.NewRunnerClient(conn)
	}
	return MyApp
}
```
