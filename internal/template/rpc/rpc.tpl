package rpc

import (
	"context"
	"time"

	"github.com/google/uuid"
	"%goMod/gen/golang/basic"
	"%goMod/gen/golang/hello"
)

type Service struct {
	Foo
}

type Foo struct {
}

var uid = uuid.NewString()

// Foo
func (s *Foo) Hello(context.Context, *basic.Empty) (*hello.Response, error) {
	time.Sleep(time.Second)
	return &hello.Response{Data: "你好世界,此消息来自机器:" + uid}, nil
}
