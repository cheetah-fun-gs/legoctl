package generated

import (
	"context"
    {{range .Imports}}handler{{ . }}
	{{end}}
	"github.com/cheetah-fun-gs/goso/pkg/handler"
)

// Handlers gnet Handlers
var Handlers = []*handler.Handler{
	{{range .Handlers}}&handler.Handler{
		Name:    "{{ .PackageTitle }}{{ .HandlerName }}",
		Nets:    handler{{ .PackageName }}.{{ .HandlerName }}NetTypes,
		Routers: handler{{ .PackageName }}.{{ .HandlerName }}Routers,
		Req:     &handler{{ .PackageName }}.{{ .HandlerName }}Req{},
		Resp:    &handler{{ .PackageName }}.{{ .HandlerName }}Resp{},
		Func: func(ctx context.Context, req, resp interface{}) error {
			return handler{{ .PackageName }}.{{ .HandlerName }}Handle(ctx, req.(*handler{{ .PackageName }}.{{ .HandlerName }}Req), resp.(*handler{{ .PackageName }}.{{ .HandlerName }}Resp))
		},
	},
    {{end}}
}
