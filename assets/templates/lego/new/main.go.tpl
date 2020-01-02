package main

import (
	// 顺序不能错
	_ "{{.ProjectName}}/cmd"
	_ "{{.ProjectName}}/internal/common"
	_ "{{.ProjectName}}/internal/svc"
)

func main() {
}
