package dev

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	filepathplus "github.com/cheetah-fun-gs/goplus/path/filepath"
)

var (
	gp       = os.Getenv("GOPATH")
	srcDir   = filepath.Join(gp, "src", "github.com", "cheetah-fun-gs", "goso")
	dstDir   = filepath.Join(gp, "src", "github.com", "cheetah-fun-gs", "goso-cli", "assets", "tpl", "init")
	replaces = []*filepathplus.ReplaceOption{
		&filepathplus.ReplaceOption{
			Old: `"github.com/cheetah-fun-gs/goso/internal`,
			New: `"{{.ProjectName}}/internal`,
		},
		&filepathplus.ReplaceOption{
			Old: `"github.com/cheetah-fun-gs/goso/cmd"`,
			New: `"{{.ProjectName}}/cmd"`,
		},
	}
)

func isSkipCopy(path string) bool {
	m := map[string]bool{
		"go.mod":    true,
		"go.sum":    true,
		"LICENSE":   true,
		"README.md": true,
	}
	if _, ok := m[filepath.Base(path)]; ok {
		return true
	}

	for _, s := range []string{
		strings.Join([]string{string(os.PathSeparator), "goso", string(os.PathSeparator), "docs", string(os.PathSeparator)}, ""),
		strings.Join([]string{string(os.PathSeparator), "goso", string(os.PathSeparator), "pkg", string(os.PathSeparator)}, ""),
	} {
		if strings.Contains(path, s) {
			return true
		}
	}
	return false
}

// Tpl 构建模板
func Tpl() {
	if err := buildTpl(srcDir, dstDir, replaces, isSkipCopy); err != nil {
		panic(err)
	}
	fmt.Printf("%s => %s success!\n", srcDir, dstDir)
}

func buildTpl(srcDir, dstDir string, replaces []*filepathplus.ReplaceOption, isSkipCopy func(name string) bool) error {
	files, err := filepathplus.Files(srcDir)
	if err != nil {
		return nil
	}

	for _, srcPath := range files {
		if isSkipCopy(srcPath) {
			continue
		}

		dstPath := strings.Replace(srcPath, srcDir, dstDir, 1) + ".tpl"

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

		if err := filepathplus.CopyFileAndReplace(srcPath, dstPath, replaces); err != nil {
			return err
		}
	}
	return nil
}
