package internal

import (
	"context"
	"github.com/xhyonline/xutil/logger"
	"net/http"
)

type HTTPServer struct {
	*http.Server
}

func (s *HTTPServer) GracefulClose() {
	if err := s.Shutdown(context.Background()); err != nil {
		logger.Errorf("HTTP 服务优雅退出失败 %s", err)
		return
	}
	logger.Info("HTTP 服务已优雅退出")
}

// Run 启动
func (s *HTTPServer) Run() {
	logger.Info("HTTP 服务成功启动")
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("HTTP 服务启动出错 %s", err)
	}
}
