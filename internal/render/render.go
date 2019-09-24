package render

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"

	filepathplus "github.com/cheetah-fun-gs/goplus/path/filepath"
)

// FileToFile 渲染一个文件
func FileToFile(data interface{}, tplFilePath, goFilePath string) error {
	tmpl, err := template.ParseFiles(tplFilePath)
	if err != nil {
		return err
	}

	isExists, err := filepathplus.Exists(goFilePath)
	if err != nil {
		return err
	}
	if isExists {
		if err := os.Remove(goFilePath); err != nil {
			return err
		}
	}

	fileObj, err := os.OpenFile(goFilePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	if err := tmpl.Execute(fileObj, data); err != nil {
		return err
	}
	return nil
}

// DirToDir 渲染一个目录
func DirToDir(data interface{}, tplFileDir, goFileDir string) error {
	files, err := filepathplus.Files(tplFileDir)
	if err != nil {
		return nil
	}

	for _, filePath := range files {
		dstPath := strings.Replace(filePath, tplFileDir, goFileDir, 1)

		dstDir := filepath.Dir(dstPath)
		isExists, err := filepathplus.Exists(dstDir)
		if err != nil {
			return err
		}
		if !isExists {
			if err := os.MkdirAll(dstDir, os.ModePerm); err != nil {
				return err
			}
		}

		if strings.HasSuffix(filePath, ".tpl") {
			dstPath = filepathplus.NoExt(dstPath)
			if err := FileToFile(data, filePath, dstPath); err != nil {
				return err
			}
		} else {
			if err := filepathplus.CopyFile(filePath, dstPath); err != nil {
				return err
			}
		}
	}
	return nil
}
