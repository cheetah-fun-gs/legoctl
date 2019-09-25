package render

import (
	"os"
	"path/filepath"
	"text/template"

	filepathplus "github.com/cheetah-fun-gs/goplus/path/filepath"
)

// File 渲染一个文件
func File(data interface{}, tplFilePath, dstFilePath string) error {
	tmpl, err := template.ParseFiles(tplFilePath)
	if err != nil {
		return err
	}

	// 存在父目录则创建父目录
	dstFileDir := filepath.Dir(dstFilePath)
	isExists, err := filepathplus.Exists(dstFileDir)
	if err != nil {
		return err
	}
	if !isExists {
		if err := os.MkdirAll(dstFileDir, os.ModePerm); err != nil {
			return err
		}
	}

	// 存在go文件则删除go文件
	isExists, err = filepathplus.Exists(dstFilePath)
	if err != nil {
		return err
	}
	if isExists {
		if err := os.Remove(dstFilePath); err != nil {
			return err
		}
	}

	fileObj, err := os.OpenFile(dstFilePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	if err := tmpl.Execute(fileObj, data); err != nil {
		return err
	}
	return nil
}
