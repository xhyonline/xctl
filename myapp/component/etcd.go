package component

import (
	"fmt"
	"github.com/xhyonline/xutil/logger"

	"github.com/xhyonline/myapp/configs"

	"github.com/xhyonline/xutil/etcd"
)

func RegisterETCD() Option {
	return func() {
		address := []string{
			fmt.Sprintf("%s:%d", configs.Instance.ETCD.Host, configs.Instance.ETCD.Port),
		}
		client, err := etcd.New(address...)
		if err != nil {
			logger.Fatalf("etcd 启动失败,地址:%s %s", address, err)
		}
		Instance.ETCD = client
	}
}
