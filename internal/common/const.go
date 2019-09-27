package common

import (
	"os"
	"path/filepath"
)

// 常用变量
var (
	GoPath       = os.Getenv("GOPATH")
	DemoRoot     = filepath.Join(GoPath, "src", "github.com", "cheetah-fun-gs", "goso")
	TplRoot      = filepath.Join(GoPath, "src", "github.com", "cheetah-fun-gs", "goso-cli", "assets", "tpl")
	TplInit      = filepath.Join(TplRoot, "init")
	TplGenerated = filepath.Join(TplRoot, "generated")
)
