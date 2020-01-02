package dev

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	filepathplus "github.com/cheetah-fun-gs/goplus/path/filepath"
	"github.com/cheetah-fun-gs/legoctl/internal/common"
)

var (
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
		"go.mod":  true,
		"go.sum":  true,
		"LICENSE": true,
		"goso":    true,
	}
	if _, ok := m[filepath.Base(path)]; ok {
		return true
	}

	for _, s := range []string{
		strings.Join([]string{string(os.PathSeparator), "goso", string(os.PathSeparator), ".git", string(os.PathSeparator)}, ""),
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
	isExists, err := filepathplus.Exists(common.TplInit)
	if err != nil {
		panic(err)
	}
	if isExists {
		if err := os.RemoveAll(common.TplInit); err != nil {
			panic(err)
		}
	}

	if err := buildTpl(common.DemoRoot, common.TplInit, replaces, isSkipCopy); err != nil {
		panic(err)
	}

	if err := filepathplus.CopyFile(filepath.Join(common.TplRoot, "README.md.tpl"), filepath.Join(common.TplInit, "README.md.tpl")); err != nil {
		panic(err)
	}
	fmt.Printf("%s => %s success!\n", common.DemoRoot, common.TplInit)
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
