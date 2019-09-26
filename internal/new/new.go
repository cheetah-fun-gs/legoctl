package new

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	execplus "github.com/cheetah-fun-gs/goplus/exec"
	filepathplus "github.com/cheetah-fun-gs/goplus/path/filepath"
	"github.com/cheetah-fun-gs/goso-cli/internal/common"
	"github.com/cheetah-fun-gs/goso-cli/pkg/render"
)

// Data 新工程的data
type Data struct {
	ProjectName string
}

// New 渲染一个目录
func New(projectPath string) {
	isExists, err := filepathplus.Exists(filepath.Join(projectPath, "main.go"))
	if err != nil {
		panic(err)
	}
	if isExists {
		fmt.Printf("new %s fail: dir is exists!\n", projectPath)
		return
	}

	if err := runNew(projectPath); err != nil {
		panic(err)
	}
	fmt.Printf("new %s success!\n", projectPath)
}

func runNew(projectPath string) error {
	projectName := filepath.Base(projectPath)

	data := &Data{ProjectName: projectName}

	files, err := filepathplus.Files(common.TplInit)
	if err != nil {
		return nil
	}

	for _, tplFilePath := range files {
		dstFilePath := filepathplus.NoExt(strings.Replace(tplFilePath, common.TplInit, projectPath, 1))
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

	os.Chdir(projectPath)
	_, stdout, err := execplus.Command("go", "mod", "init", projectName)
	if err != nil {
		return err
	}
	for _, line := range stdout {
		fmt.Printf(line)
	}
	_, stdout, err = execplus.Command("go", "mod", "vendor")
	if err != nil {
		return err
	}
	for _, line := range stdout {
		fmt.Printf(line)
	}
	return nil
}
