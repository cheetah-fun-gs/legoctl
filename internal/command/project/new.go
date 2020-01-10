package project

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	execplus "github.com/cheetah-fun-gs/goplus/exec"
	filepathplus "github.com/cheetah-fun-gs/goplus/path/filepath"
	"github.com/cheetah-fun-gs/legoctl/internal/common"
	"github.com/cheetah-fun-gs/legoctl/pkg/render"
)

// Data 新工程的data
type newData struct {
	PackageName string
}

// New 渲染一个目录
func New(projectPath string, opt *Opt) error {
	projectPath = common.GetProjectPath(projectPath)
	if opt.PackageName == "" {
		opt.PackageName = filepath.Base(projectPath)
	}

	isExists, err := filepathplus.Exists(filepath.Join(projectPath, "main.go"))
	if err != nil {
		return err
	}
	if isExists {
		return fmt.Errorf("%v is exists", projectPath)
	}

	// 确认模板目录
	templateNewPath := filepath.Join(common.GetTemplateRoot(opt.TemplateName), "new")
	isExists, err = filepathplus.Exists(templateNewPath)
	if err != nil {
		return err
	}
	if !isExists {
		return fmt.Errorf("%v is not exists", templateNewPath)
	}

	if err := buildNew(templateNewPath, projectPath, opt); err != nil {
		return err
	}
	fmt.Printf("new %s success!\n", projectPath)
	return nil
}

func buildNew(templateNewPath, projectPath string, opt *Opt) error {
	data := &newData{PackageName: opt.PackageName}

	files, err := filepathplus.Files(templateNewPath)
	if err != nil {
		return nil
	}

	for _, tplFilePath := range files {
		dstFilePath := filepathplus.NoExt(strings.Replace(tplFilePath, templateNewPath, projectPath, 1))
		if err := render.File(data, tplFilePath, dstFilePath); err != nil {
			return err
		}

		if strings.HasSuffix(dstFilePath, ".go") {
			_, stdout, err := execplus.Command("gofmt", "-w", dstFilePath)
			for _, line := range stdout {
				fmt.Printf(line)
			}
			if err != nil {
				return err
			}
		}
	}

	os.Chdir(projectPath)
	_, stdout, err := execplus.Command("go", "mod", "init", opt.PackageName)
	for _, line := range stdout {
		fmt.Printf(line)
	}
	if err != nil {
		return err
	}

	_, stdout, err = execplus.Command("go", "mod", "vendor")
	for _, line := range stdout {
		fmt.Printf(line)
	}
	if err != nil {
		return err
	}
	return nil
}
