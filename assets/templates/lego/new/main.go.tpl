package main

import (
	// 顺序不能错
	_ "{{.PackageName}}/cmd"
	_ "{{.PackageName}}/internal/common"
	_ "{{.PackageName}}/internal/svc"
)

//go:generate legoctl project gen $GOFILE
func main() {
}
