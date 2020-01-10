// Code generated  DO NOT EDIT.
package handler

import (
	"context"
	"fmt"

	jsonplus "github.com/cheetah-fun-gs/goplus/encoding/json"
	mlogger "github.com/cheetah-fun-gs/goplus/multier/multilogger"
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
	beforeInjectFunc []func(ctx context.Context, req handler{{ .PackageName }}.{{ .HandlerName }}Req)
	behindInjectFunc []func(ctx context.Context, req handler{{ .PackageName }}.{{ .HandlerName }}Req, resp handler{{ .PackageName }}.{{ .HandlerName }}Resp)
}

// InjectBeforeFunc 注入前置函数
func (h *{{ .PackageTitle }}{{ .HandlerName }}Handler) InjectBeforeFunc(f ...func(ctx context.Context, req handler{{ .PackageName }}.{{ .HandlerName }}Req)) {
	h.beforeInjectFunc = append(h.beforeInjectFunc, f...)
}

// InjectBehindFunc 注入后置函数
func (h *{{ .PackageTitle }}{{ .HandlerName }}Handler) InjectBehindFunc(f ...func(ctx context.Context, req handler{{ .PackageName }}.{{ .HandlerName }}Req, resp handler{{ .PackageName }}.{{ .HandlerName }}Resp)) {
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

	// 前置注入
	if len(h.beforeInjectFunc) > 0 {
		beforeInjectReq := handler{{ .PackageName }}.{{ .HandlerName }}Req{}
		if err := jsonplus.Convert(reqBody, &beforeInjectReq); err != nil {
			mlogger.Warnc(ctx, "jsonplus.Convert err: %v", err)
		} else {
			for _, f := range h.beforeInjectFunc {
				go f(ctx, beforeInjectReq)
			}
		}
	}

	// 真正处理
	if err := handler{{ .PackageName }}.{{ .HandlerName }}Handle(ctx, reqBody, respBody); err != nil {
		return err
	}

	// 后置注入
	if len(h.behindInjectFunc) > 0 {
		behindInjectReq := handler{{ .PackageName }}.{{ .HandlerName }}Req{}
		behindInjectResp := handler{{ .PackageName }}.{{ .HandlerName }}Resp{}
		errReq := jsonplus.Convert(reqBody, &behindInjectReq)
		errResp := jsonplus.Convert(respBody, &behindInjectResp)
		if errReq != nil || errResp != nil {
			mlogger.Warnc(ctx, "jsonplus.Convert errReq: %v, errResp: %v", errReq, errResp)
		} else {
			for _, f := range h.behindInjectFunc {
				go f(ctx, behindInjectReq, behindInjectResp)
			}
		}
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
var Handlers = map[string]legocore.Handler{
	{{range .Handlers}}"{{ .PackageTitle }}{{ .HandlerName }}" : {{ .PackageTitle }}{{ .HandlerName }},
    {{end}}
}
