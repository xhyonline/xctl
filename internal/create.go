package internal

import (
	"embed"
	_ "embed"
	"os"
)

//go:embed template/*
var FS embed.FS

var currentPath, _ = os.Getwd()

// CreateProject 创建项目
func CreateProject(args *CreateArgs) {
	currentPath = currentPath + "/" + args.AppName
	_ = os.MkdirAll(currentPath, 777)
	args.createFile("template", currentPath)
}
