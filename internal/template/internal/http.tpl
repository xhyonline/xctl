package internal

import (
	"context"
	. "%goMod/component" // 忽略包名
	"net/http"
	"os"
)

type HTTPServer struct {
	*http.Server
}

func (s *HTTPServer) GracefulClose() {
	if err := s.Shutdown(context.Background()); err != nil {
		Logger.Error("HTTP 服务优雅退出失败", err)
	}
	Logger.Info("HTTP 服务已优雅退出")
}

// Run 启动
func (s *HTTPServer) Run() {
	Logger.Info("HTTP 服务成功启动")
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		Logger.Error("HTTP 服务启动出错", err)
		os.Exit(1)
	}
}
