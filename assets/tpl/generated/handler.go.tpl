package generated

import (
	"context"
    {{range .Imports}}handler{{ . }}
	{{end}}
	"github.com/cheetah-fun-gs/goso/pkg/handler"
	sohttp "github.com/cheetah-fun-gs/goso/pkg/net/sohttp"
)

// Handlers gnet Handlers
var Handlers = []*handler.Handler{
	{{range .Handlers}}&handler.Handler{
		Name:    "{{ .PackageTitle }}{{ .HandlerName }}",
		Nets:    handlers{{ .PackageName }}.{{ .HandlerName }}NetTypes,
		Routers: sohttp.NewRouters(handlers{{ .PackageName }}.{{ .HandlerName }}URIS, handlers{{ .PackageName }}.{{ .HandlerName }}HTTPMethods),
		Req:     &handlers{{ .PackageName }}.{{ .HandlerName }}Req{},
		Resp:    &handlers{{ .PackageName }}.{{ .HandlerName }}Resp{},
		Func: func(ctx context.Context, req, resp interface{}) error {
			return handlers{{ .PackageName }}.{{ .HandlerName }}Handle(ctx, req.(*handlers{{ .PackageName }}.{{ .HandlerName }}Req), resp.(*handlers{{ .PackageName }}.{{ .HandlerName }}Resp))
		},
	},
    {{end}}
}
