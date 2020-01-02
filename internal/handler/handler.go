package handler

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/cheetah-fun-gs/goplus/gostyle"
	filepathplus "github.com/cheetah-fun-gs/goplus/path/filepath"
)

// FileRender 处理器文件渲染对象
type FileRender struct {
	Imports  []string // handlersxxx "%s/internal/biz/handler/xxx"
	Handlers []*Render
}

// Render 处理器渲染对象
type Render struct {
	PackageTitle string // 首字母大写
	PackageName  string // 全小写
	HandlerName  string
}

var handlerChecks = map[string]ast.ObjKind{
	"Req":  ast.Typ,
	"Resp": ast.Typ,
}

func findHandlerRender(packageName string, objs []*ast.Object) (bool, *Render) {
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

	return true, &Render{
		PackageTitle: gostyle.FormatToCamelCase(packageName),
		PackageName:  packageName,
		HandlerName:  handlerName,
	}
}

// Scan 扫描
func Scan(projectPath string) (*FileRender, error) {
	goFilePaths, err := filepathplus.Files(filepath.Join(projectPath, "internal", "biz", "handler"))
	if err != nil {
		return nil, err
	}

	projectName := filepath.Base(projectPath)

	var imports = map[string]string{}
	var handlers = []*Render{}

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

	handleFileRender := &FileRender{
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
