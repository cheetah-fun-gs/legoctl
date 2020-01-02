package common

import (
	"fmt"
	"os"
	"path/filepath"

	filepathplus "github.com/cheetah-fun-gs/goplus/path/filepath"
)

// GetCtlRoot 获取工具的根路径
func GetCtlRoot() string {
	// 确认是否工作目录
	workFolder, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	if filepath.Base(workFolder) == "legoctl" {
		return workFolder
	}
	// 确认是否gopath
	gopathFolder := filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "cheetah-fun-gs", "legoctl")
	ok, err := filepathplus.Exists(gopathFolder)
	if err != nil {
		panic(err)
	}
	if ok {
		return gopathFolder
	}
	panic("not found legoctl's root")
}

// GetTemplateRoot 获取模板根目录
func GetTemplateRoot(templateName string) string {
	return filepath.Join(GetCtlRoot(), "assets", "templates", templateName)
}

// GetProjectPath 获取目标项目地址
func GetProjectPath(path string) string {
	info, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	if info.IsDir() {
		return path
	}
	if filepath.Base(path) == "main.go" {
		return filepath.Dir(path)
	}
	panic(fmt.Errorf("%v is not projectPath", path))
}
