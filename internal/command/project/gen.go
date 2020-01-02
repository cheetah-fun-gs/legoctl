package project

import (
	"fmt"
	"path/filepath"
	"strings"

	execplus "github.com/cheetah-fun-gs/goplus/exec"
	filepathplus "github.com/cheetah-fun-gs/goplus/path/filepath"
	"github.com/cheetah-fun-gs/legoctl/internal/common"
	"github.com/cheetah-fun-gs/legoctl/internal/handler"
	"github.com/cheetah-fun-gs/legoctl/pkg/render"
)

// Gen 生产项目代码
func Gen(projectPath, templateName string) error {
	projectPath = common.GetProjectPath(projectPath)
	// 确认模板目录
	templateGenPath := filepath.Join(common.GetTemplateRoot(templateName), "gen")
	isExists, err := filepathplus.Exists(templateGenPath)
	if err != nil {
		return err
	}
	if !isExists {
		return fmt.Errorf("%v is not exists", templateGenPath)
	}

	handlerFileRender, err := handler.Scan(projectPath)
	if err != nil {
		return err
	}

	projectGenPath := filepath.Join(projectPath, "internal", "generated")

	if err = buildGen(handlerFileRender, templateGenPath, projectGenPath); err != nil {
		return err
	}

	fmt.Printf("gen %s success!\n", projectPath)
	return nil
}

func buildGen(data *handler.FileRender, templateGenPath, projectGenPath string) error {
	files, err := filepathplus.Files(templateGenPath)
	if err != nil {
		return nil
	}

	for _, tplFilePath := range files {
		dstFilePath := filepathplus.NoExt(strings.Replace(tplFilePath, templateGenPath, projectGenPath, 1))
		if err := render.File(data, tplFilePath, dstFilePath); err != nil {
			return err
		}

		if strings.HasSuffix(dstFilePath, ".go") {
			_, stdout, err := execplus.Command("gofmt", "-w", dstFilePath)
			if err != nil {
				return err
			}
			for _, line := range stdout {
				fmt.Printf(line)
			}
		}
	}
	return nil
}
