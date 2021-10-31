package internal

import (
	"os"
	"path"
	"strings"

	"github.com/xhyonline/xutil/helper"
)

// CreateArgs 接收参数
type CreateArgs struct {
	WithMySQL, WithRedis, WithEtcd,
	WithGithubAction, WithHTTPServer,
	WithGRPCServer bool
	AppName, Mod string
}

// replaceContentByTag 通过标签更换内容
func (s *CreateArgs) replaceContentByTag(input string) string {
	register, server, imports := s.getComponent()
	var label = map[string]string{
		"%goMod":             s.Mod,
		"%configs":           s.getWithConfig(),
		"%componentRegister": register,
		"%componentSever":    server,
		"%componentImport":   imports,
		"%appName":           s.AppName,
	}
	for k, v := range label {
		input = strings.ReplaceAll(input, k, v)
	}
	return input
}

func (s *CreateArgs) getWithConfig() string {
	var tmp = make([]string, 0)
	if s.WithMySQL {
		tmp = append(tmp, "configs.WithMySQL()")
	}
	if s.WithRedis {
		tmp = append(tmp, "configs.WithRedis()")
	}
	if s.WithGRPCServer {
		tmp = append(tmp, "configs.WithETCD()")
	}
	if len(tmp) == 0 {
		return ""
	}
	return strings.Join(tmp, ",")
}

func (s *CreateArgs) getComponent() (register, server, imports string) {
	var registerComponent = make([]string, 0)
	if s.WithMySQL {
		registerComponent = append(registerComponent, "component.RegisterMySQL()")
		server += "MySQL *gorm.DB\n"
		imports += `"gorm.io/gorm"` + "\n"
	}
	if s.WithRedis {
		registerComponent = append(registerComponent, "component.RegisterRedis()")
		server += "Redis *kv.RClient\n"
		imports += `"github.com/xhyonline/xutil/kv"` + "\n"
	}
	if s.WithGRPCServer {
		registerComponent = append(registerComponent, "component.RegisterETCD()")
		server += "ETCD  *clientv3.Client\n"
		imports += `"go.etcd.io/etcd/clientv3"`
	}
	if len(registerComponent) == 0 {
		return "", "", ""
	}
	return strings.Join(registerComponent, ","), strings.TrimRight(server, "\n"), imports
}

// createFile 递归拷贝模板,并且格式化创建 Go 创建文件
func (s *CreateArgs) createFile(tplPath, filePath string) {
	d, _ := FS.ReadDir(tplPath)
	for _, item := range d {
		createPath := filePath + "/" + item.Name()
		tpl := tplPath + "/" + item.Name()
		if s.skipCreateFile(createPath) {
			continue
		}
		if item.IsDir() {
			_ = os.MkdirAll(createPath, 777)
			s.createFile(tpl, createPath)
			continue
		}
		body, _ := FS.ReadFile(tpl)
		ext := path.Ext(createPath)
		switch ext {
		case ".tpl_remove":
			createPath = strings.Replace(createPath, ext, "", 1)
		case ".tpl":
			createPath = strings.Replace(createPath, ext, ".go", 1)
		case ".maintpl":
			createPath = strings.Replace(createPath, "http.maintpl", "main.go", 1)
			createPath = strings.Replace(createPath, "server.maintpl", "main.go", 1)
		}
		// 替换标签
		input := s.replaceContentByTag(string(body))
		_ = helper.FilePutContents(createPath, input, helper.ContentCover)
	}
}

// skipCreateFile 是否跳过该文件的创建
func (s *CreateArgs) skipCreateFile(path string) bool {
	if !s.WithMySQL && helper.InArray(path, []string{
		currentPath + "/component/mysql.tpl",
	}) {
		return true
	}
	if !s.WithRedis && helper.InArray(path, []string{
		currentPath + "/component/redis.tpl",
	}) {
		return true
	}
	if !s.WithGithubAction && helper.InArray(path, []string{
		currentPath + "/.github",
	}) {
		return true
	}
	// 纯 HTTP 服务,不带 GRPC Client
	if s.WithHTTPServer && !s.WithGRPCServer && helper.InArray(path, []string{
		currentPath + "/gen",
		currentPath + "/rpc",
		currentPath + "/protobuf",
		currentPath + "/component/etcd.tpl",
		currentPath + "/configs/common/etcd.toml",
		currentPath + "/server.maintpl",
	}) {
		return true
	}
	// HTTP 服务带上 GRPC Client
	if s.WithHTTPServer && s.WithGRPCServer && helper.InArray(path, []string{
		currentPath + "/gen",
		currentPath + "/protobuf",
		currentPath + "/rpc/rpc.tpl",
		currentPath + "/server.maintpl",
	}) {
		return true
	}
	// 纯 GRPC 服务
	if s.WithGRPCServer && !s.WithHTTPServer && helper.InArray(path, []string{
		currentPath + "/internal/http.tpl",
		currentPath + "/middleware",
		currentPath + "/router",
		currentPath + "/http.maintpl",
	}) {
		return true
	}
	return false
}
