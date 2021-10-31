package internal

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/xhyonline/xutil/helper"

	"github.com/xhyonline/xutil/shell"
)

// GeneratePb
func GeneratePb() error {
	// 插件检查
	if _, err := exec.LookPath("protoc"); err != nil {
		return err
	}
	if _, err := exec.LookPath("protoc-gen-gogofaster"); err != nil {
		return err
	}
	if exists, _ := helper.PathExists("./protobuf"); !exists {
		return fmt.Errorf("当前路径下不存在 protobuf 目录")
	}
	// 清理
	_, _ = shell.Command("bash", "-c", "rm -rf ./gen/*")

	command := "protoc --gogofaster_out=plugins=grpc:./gen"
	result, err := shell.Command("bash", "-c", fmt.Sprintf("%s %s",
		command, "./protobuf/*.proto"))
	if err != nil {
		return fmt.Errorf("%s", result)
	}
	files, err := ioutil.ReadDir("./protobuf")
	if err != nil {
		return err
	}
	for _, f := range files {
		if !f.IsDir() {
			continue
		}
		result, err = shell.Command("bash", "-c", fmt.Sprintf("%s %s",
			command, "./protobuf/"+f.Name()+"/*.proto"))
		if err != nil {
			return fmt.Errorf("%s", result)
		}
	}
	p := "./gen"
	var flag bool
	var firstDirName string
FOR:
	for {
		files, err = ioutil.ReadDir(p)
		if err != nil {
			return err
		}
		for _, f := range files {
			if f.IsDir() && f.Name() == "gen" {
				p += "/gen"
				break FOR
			}
			if f.IsDir() {
				if !flag {
					firstDirName = f.Name()
					flag = true
				}
				p += "/" + f.Name()
				continue FOR
			}
		}
		// 没有找到 gen 目录,直接退出,并且删除刚才生成的 pb 文件
		_, _ = shell.Command("bash", "-c", "rm -rf ./gen/*")
		return fmt.Errorf("protobuf 中, go_package 未定义 gen 目录")
	}
	_, _ = shell.Command("bash", "-c", "cp -rf "+p+" ./")
	_, _ = shell.Command("bash", "-c", fmt.Sprintf("rm -rf ./gen/%s", firstDirName))
	return nil
}
