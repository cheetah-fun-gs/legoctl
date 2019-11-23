// Code generated  DO NOT EDIT.
package generated

import (
	"context"
    {{range .Imports}}handler{{ . }}
	{{end}}
	"github.com/cheetah-fun-gs/goso/pkg/handler"
	"github.com/cheetah-fun-gs/goso/pkg/so"
)

{{range .Handlers}}// {{ .PackageTitle }}{{ .HandlerName }}Handler {{ .PackageTitle }}{{ .HandlerName }}Handler
type {{ .PackageTitle }}{{ .HandlerName }}Handler struct {
	*handler.Handler
}

// CloneReq CloneReq
func (h *{{ .PackageTitle }}{{ .HandlerName }}Handler) CloneReq() interface{} {
	return &handler{{ .PackageName }}.{{ .HandlerName }}Req{}
}

// CloneResp CloneResp
func (h *{{ .PackageTitle }}{{ .HandlerName }}Handler) CloneResp() interface{} {
	return &handler{{ .PackageName }}.{{ .HandlerName }}Resp{}
}
{{end}}

// Handlers Handlers
var Handlers = []so.Handler{
	{{range .Handlers}}&{{ .PackageTitle }}{{ .HandlerName }}Handler{
		Handler: &handler.Handler{
			Name:    "{{ .PackageTitle }}{{ .HandlerName }}",
			Nets:    handler{{ .PackageName }}.{{ .HandlerName }}NetTypes,
			Routers: handler{{ .PackageName }}.{{ .HandlerName }}Routers,
			Func: func(ctx context.Context, req, resp interface{}) error {
				return handler{{ .PackageName }}.{{ .HandlerName }}Handle(ctx, req.(*handler{{ .PackageName }}.{{ .HandlerName }}Req), resp.(*handler{{ .PackageName }}.{{ .HandlerName }}Resp))
			},
		},
	},
    {{end}}
}
