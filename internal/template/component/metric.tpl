package component

import (
	"%goMod/configs"
	"github.com/xhyonline/xutil/helper"
	"github.com/xhyonline/xutil/logger"
	"github.com/xhyonline/xutil/metrics"
	"net"
	"net/http"
)

// WithPProf 监控
func WithPProf() Option {
	return func() {
		go func() {
			if err := http.ListenAndServe(internalIP()+":0", nil); err != nil {
				logger.Fatalf("pprof 服务启动失败")
			}
		}()
	}
}

// WithPrometheus 监控
func WithPrometheus() Option {
	return func() {
		metrics.Init(configs.Instance.PrometheusGateWay.Host+":"+configs.Instance.PrometheusGateWay.Port, configs.Name)
	}
}

// internalIP 获取内网 IP
func internalIP() string {
	var address = "127.0.0.1"
	addr, err := helper.IntranetAddress()
	if err != nil {
		logger.Fatalf("获取内网地址失败,服务停止 %s", err)
	}
	v := addr["eth0"]
	var ip net.IP
	for _, item := range v {
		if ip = item.To4(); ip != nil {
			break
		}
	}
	if ip != nil {
		address = ip.String()
	} else {
		logger.Errorf("未发现 IPv4 地址,将使用 %s 替代", address)
	}

	return address
}
