package internal

import (
	"github.com/xhyonline/xutil/helper"
	"os"
	"path"
	"strings"
)

// Args 接收参数
type Args struct {
	WithMySQL, WithRedis, WithEtcd,
	WithGithubAction, WithHTTPServer bool
	AppName, Mod string
}

// replaceContentByTag 通过标签更换内容
func (s *Args) replaceContentByTag(input string) string {
	register, server, imports := s.getComponent()
	var label = map[string]string{
		"%goMod":             s.Mod,
		"%configs":           s.getWithConfig(),
		"%componentRegister": register,
		"%componentSever":    server,
		"%componentImport":   imports,
	}
	for k, v := range label {
		input = strings.ReplaceAll(input, k, v)
	}
	return input
}

func (s *Args) getWithConfig() string {
	var tmp = make([]string, 0)
	if s.WithMySQL {
		tmp = append(tmp, "configs.WithMySQL()")
	}
	if s.WithRedis {
		tmp = append(tmp, "configs.WithRedis()")
	}
	if len(tmp) == 0 {
		return ""
	}
	return strings.Join(tmp, ",")
}

func (s *Args) getComponent() (register, server, imports string) {
	var registerComponent = make([]string, 0)
	if s.WithMySQL {
		registerComponent = append(registerComponent, "RegisterMySQL()")
		server += "MySQL *gorm.DB\n"
		imports += `"gorm.io/gorm"` + "\n"
	}
	if s.WithRedis {
		registerComponent = append(registerComponent, "RegisterRedis()")
		server += "Redis *kv.RClient\n"
		imports += `"github.com/xhyonline/xutil/kv"` + "\n"
	}
	if len(registerComponent) == 0 {
		return "", "", ""
	}
	return "," + strings.Join(registerComponent, ","), strings.TrimRight(server, "\n"), imports
}

// createFile 递归拷贝模板,并且格式化创建 Go 创建文件
func (s *Args) createFile(tplPath, filePath string) {
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
		}
		// 替换标签
		input := s.replaceContentByTag(string(body))
		_ = helper.FilePutContents(createPath, input, helper.ContentCover)
	}
}

// skipCreateFile 是否跳过该文件的创建
func (s *Args) skipCreateFile(path string) bool {
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
	return false
}
