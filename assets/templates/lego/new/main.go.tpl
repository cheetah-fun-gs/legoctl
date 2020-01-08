package main

import (
	// 顺序不能错
	_ "{{.PackageName}}/cmd/web"
	_ "{{.PackageName}}/internal/svc"
)

//go:generate legoctl project gen $GOFILE -p github.com/cheetah-fun-gs/lego
func main() {
}
