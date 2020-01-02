// Code generated  DO NOT EDIT.
package handler

import (
	"context"
	"fmt"

    {{range .Imports}}handler{{ . }}
	{{end}}
	legocore "github.com/cheetah-fun-gs/lego/pkg/core"
)

{{range .Handlers}}// {{ .PackageTitle }}{{ .HandlerName }}Handler {{ .PackageTitle }}{{ .HandlerName }}Handler
type {{ .PackageTitle }}{{ .HandlerName }}Handler struct {
	*legocore.DefaultHandler
	SvcNames         []string
	HTTPMethods      []string
	HTTPPaths        []string
	beforeInjectFunc []func(ctx context.Context, req *handler{{ .PackageName }}.{{ .HandlerName }}Req)
	behindInjectFunc []func(ctx context.Context, req *handler{{ .PackageName }}.{{ .HandlerName }}Req, resp *handler{{ .PackageName }}.{{ .HandlerName }}Resp)
}

// InjectBeforeFunc 注入前置函数
func (h *{{ .PackageTitle }}{{ .HandlerName }}Handler) InjectBeforeFunc(f ...func(ctx context.Context, req *handler{{ .PackageName }}.{{ .HandlerName }}Req)) {
	h.beforeInjectFunc = append(h.beforeInjectFunc, f...)
}

// InjectBehindFunc 注入后置函数
func (h *{{ .PackageTitle }}{{ .HandlerName }}Handler) InjectBehindFunc(f ...func(ctx context.Context, req *handler{{ .PackageName }}.{{ .HandlerName }}Req, resp *handler{{ .PackageName }}.{{ .HandlerName }}Resp)) {
	h.behindInjectFunc = append(h.behindInjectFunc, f...)
}

// CloneReq CloneReq
func (h *{{ .PackageTitle }}{{ .HandlerName }}Handler) CloneReq() interface{} {
	return &handler{{ .PackageName }}.{{ .HandlerName }}Req{}
}

// CloneResp CloneResp
func (h *{{ .PackageTitle }}{{ .HandlerName }}Handler) CloneResp() interface{} {
	return &handler{{ .PackageName }}.{{ .HandlerName }}Resp{}
}

// Handle 处理函数
func (h *{{ .PackageTitle }}{{ .HandlerName }}Handler) Handle(ctx context.Context, req, resp interface{}) error {
	reqBody, okReq := req.(*handler{{ .PackageName }}.{{ .HandlerName }}Req)
	respBody, okResp := resp.(*handler{{ .PackageName }}.{{ .HandlerName }}Resp)
	if !okReq || !okResp {
		return fmt.Errorf("req or resp type error")
	}

	for _, f := range h.beforeInjectFunc {
		go f(ctx, reqBody)
	}
	if err := handler{{ .PackageName }}.{{ .HandlerName }}Handle(ctx, reqBody, respBody); err != nil {
		return err
	}
	for _, f := range h.behindInjectFunc {
		go f(ctx, reqBody, respBody)
	}
	return nil
}

// {{ .PackageTitle }}{{ .HandlerName }} handler
var {{ .PackageTitle }}{{ .HandlerName }} = &{{ .PackageTitle }}{{ .HandlerName }}Handler{
	DefaultHandler: &legocore.DefaultHandler{
		Name: "{{ .PackageTitle }}{{ .HandlerName }}",
	},
	SvcNames:    handler{{ .PackageName }}.{{ .HandlerName }}SvcNames,
	HTTPMethods: handler{{ .PackageName }}.{{ .HandlerName }}HTTPMethods,
	HTTPPaths:   handler{{ .PackageName }}.{{ .HandlerName }}HTTPPaths,
}
{{end}}

// Handlers 所有handler
var Handlers = []legocore.Handler{
	{{range .Handlers}}{{ .PackageTitle }}{{ .HandlerName }},
    {{end}}
}
