package component

import (
	"sync"

    %componentImport
)

// Server 组件服务
type Server struct {
	%componentSever
}

var (
	Instance *Server
	once     sync.Once
)

type Option func()

// ComponentInit 初始化组建服务
func ComponentInit(options ...Option) {
	once.Do(func() {
		Instance = new(Server)
		for _, f := range options {
			f()
		}
	})
}
