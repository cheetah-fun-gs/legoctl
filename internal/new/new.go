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
func New(projectPath string) error {
	projectName := filepath.Base(projectPath)

	data := &Data{ProjectName: projectName}

	files, err := filepathplus.Files(common.TplRoot)
	if err != nil {
		return nil
	}

	for _, tplFilePath := range files {
		dstFilePath := filepathplus.NoExt(strings.Replace(tplFilePath, common.TplRoot, projectPath, 1))
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

	fmt.Printf("new %s success!\n", projectPath)

	os.Chdir(projectPath)
	_, stdout, err := execplus.Command("go", "mod")
	if err != nil {
		return err
	}
	for _, line := range stdout {
		fmt.Printf(line)
	}
	_, stdout, err = execplus.Command("go", "mod", "tidy")
	if err != nil {
		return err
	}
	for _, line := range stdout {
		fmt.Printf(line)
	}
	return nil
}
