package gen

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	execplus "github.com/cheetah-fun-gs/goplus/exec"
	"github.com/cheetah-fun-gs/goplus/gostyle"
	filepathplus "github.com/cheetah-fun-gs/goplus/path/filepath"
	"github.com/cheetah-fun-gs/legoctl/internal/common"
	"github.com/cheetah-fun-gs/legoctl/pkg/render"
)

// Gen 生产项目代码
func Gen(projectPath string) {
	handlerFileRender, err := handerFileScan(projectPath)
	if err != nil {
		panic(err)
	}

	tplFilePath := filepath.Join(common.TplGenerated, "handler.go.tpl")
	goFilePath := filepath.Join(projectPath, "internal", "generated", "handler.go")

	if err := render.File(handlerFileRender, tplFilePath, goFilePath); err != nil {
		panic(err)
	}

	_, stdout, err := execplus.Command("gofmt", "-w", goFilePath)
	if err != nil {
		panic(err)
	}
	for _, line := range stdout {
		fmt.Printf(line)
	}

	fmt.Printf("gen %s success!\n", projectPath)
}

// HandlerFileRender 处理器文件渲染对象
type HandlerFileRender struct {
	Imports  []string // handlersxxx "%s/internal/biz/handler/xxx"
	Handlers []*HandlerRender
}

// HandlerRender 处理器渲染对象
type HandlerRender struct {
	PackageTitle string // 首字母大写
	PackageName  string // 全小写
	HandlerName  string
}

var handlerChecks = map[string]ast.ObjKind{
	"NetTypes": ast.Var,
	"Routers":  ast.Var,
	"Req":      ast.Typ,
	"Resp":     ast.Typ,
}

func findHandlerRender(packageName string, objs []*ast.Object) (bool, *HandlerRender) {
	var handlerName string
	for _, o := range objs {
		if o.Kind == ast.Fun && strings.HasSuffix(o.Name, "Handle") {
			if len(o.Name) <= 6 {
				fmt.Printf("%s's name is too short, ignore it\n", o.Name)
				continue
			}
			split := strings.Split(o.Name, "")
			handlerName = strings.Join(split[:len(split)-6], "")
			break
		}
	}

	if handlerName == "" { // 没有发现 handler
		return false, nil
	}

	checked := map[string]bool{}
	for k := range handlerChecks {
		checked[k] = false
	}

	for k, v := range handlerChecks {
		for _, o := range objs {
			if o.Kind == v && o.Name == fmt.Sprintf("%s%s", handlerName, k) {
				checked[k] = true
			}
		}
	}

	isCheckPass := true
	for k, v := range checked {
		if !v {
			fmt.Printf("%sHandle has no %s%s!\n", handlerName, handlerName, k)
			isCheckPass = false
		}
	}

	if !isCheckPass {
		return false, nil
	}

	handlerRender := &HandlerRender{
		PackageTitle: gostyle.FormatToCamelCase(packageName),
		PackageName:  packageName,
		HandlerName:  handlerName,
	}
	return true, handlerRender
}

func handerFileScan(projectPath string) (*HandlerFileRender, error) {
	goFilePaths, err := filepathplus.Files(filepath.Join(projectPath, "internal", "biz", "handler"))
	if err != nil {
		return nil, err
	}

	projectName := filepath.Base(projectPath)

	var imports = map[string]string{}
	var handlers = []*HandlerRender{}

	for _, goFilePath := range goFilePaths {
		packageName, objs, err := astScan(goFilePath)
		if err != nil {
			return nil, err
		}

		ok, handlerRender := findHandlerRender(packageName, objs)
		if ok {
			if _, ok := imports[packageName]; !ok {
				imports[packageName] = fmt.Sprintf("%s/internal/biz/handler/%s", projectName, packageName)
			}
			handlers = append(handlers, handlerRender)
		}
	}

	handleFileRender := &HandlerFileRender{
		Imports:  []string{},
		Handlers: handlers,
	}
	for k, v := range imports {
		handleFileRender.Imports = append(handleFileRender.Imports,
			fmt.Sprintf(`%s "%s"`, k, v))
	}
	return handleFileRender, nil
}

func astScan(goFilePath string) (packageName string, objs []*ast.Object, err error) {
	goFileFd, err := os.Open(goFilePath)
	if err != nil {
		return
	}
	defer goFileFd.Close()

	fileData, err := ioutil.ReadAll(goFileFd)
	if err != nil {
		return
	}

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, filepath.Base(goFilePath), string(fileData), parser.ParseComments)
	if err != nil {
		return
	}

	packageName = f.Name.Name

	objs = []*ast.Object{}
	for _, o := range f.Scope.Objects {
		objs = append(objs, o)
	}
	return
}
