package template

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	filepathplus "github.com/cheetah-fun-gs/goplus/path/filepath"
	"github.com/cheetah-fun-gs/legoctl/internal/common"
)

var (
	templateReplaces = map[string][]*filepathplus.ReplaceOption{
		"lego": []*filepathplus.ReplaceOption{
			&filepathplus.ReplaceOption{
				Old: `"github.com/cheetah-fun-gs/lego/internal`,
				New: `"{{.PackageName}}/internal`,
			},
			&filepathplus.ReplaceOption{
				Old: `"github.com/cheetah-fun-gs/lego/cmd"`,
				New: `"{{.PackageName}}/cmd"`,
			},
		},
	}
)

func isSkipCopy(path, projectName string) bool {
	// 跳过的文件
	skipFileNames := []string{"go.mod", "go.sum", "LICENSE", projectName}
	for _, name := range skipFileNames {
		if filepath.Base(path) == name {
			return true
		}
	}

	// 跳过的目录
	skipDirNames := []string{".git", "docs", "pkg", "vendor"}
	for _, name := range skipDirNames {
		s := strings.Join([]string{string(os.PathSeparator), projectName, string(os.PathSeparator), name, string(os.PathSeparator)}, "")
		if strings.Contains(path, s) {
			return true
		}
	}
	return false
}

// New 创建模板
func New(projectPath string, opt *Opt) error {
	// 确认目标目录
	isExists, err := filepathplus.Exists(projectPath)
	if err != nil {
		return err
	}
	if !isExists {
		return fmt.Errorf("%v is not exist", projectPath)
	}
	// 确认模板目录
	templateNewPath := filepath.Join(common.GetTemplateRoot(opt.TemplateName), "new")
	isExists, err = filepathplus.Exists(templateNewPath)
	if err != nil {
		return err
	}
	if isExists {
		if err := os.RemoveAll(templateNewPath); err != nil {
			return err
		}
	}

	if err := buildNew(projectPath, templateNewPath, opt.TemplateName, isSkipCopy); err != nil {
		return err
	}

	fmt.Printf("%s => %s success!\n", projectPath, templateNewPath)
	return nil
}

func buildNew(srcDir, dstDir, templateName string, isSkipCopy func(path, projectPath string) bool) error {
	files, err := filepathplus.Files(srcDir)
	if err != nil {
		return nil
	}

	for _, srcPath := range files {
		if isSkipCopy(srcPath, filepath.Base(srcDir)) {
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

		if err := filepathplus.CopyFileAndReplace(srcPath, dstPath, templateReplaces[templateName]); err != nil {
			return err
		}
	}
	return nil
}
