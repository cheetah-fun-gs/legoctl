package lnet

import (
	"net/http"

	"{{.ProjectName}}/internal/biz/handlers"
	"{{.ProjectName}}/internal/generated/lnet"
	sohttp "github.com/cheetah-fun-gs/goso/pkg/net/sohttp"
	"github.com/gin-gonic/gin"
)

// SoHTTP 获取 gnet http 服务
func SoHTTP() (*sohttp.SoHTTP, error) {
	s, err := sohttp.NewLNet()
	if err != nil {
		return nil, err
	}

	// 注册首页
	s.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to goso")
	})

	if err := s.SetConfig(&sohttp.Config{Ports: []int{8600}}); err != nil {
		return nil, err
	}

	if err := s.SetErrorNetFunc(handlers.HandleCommonRespSoNet); err != nil {
		return nil, err
	}

	// 最后注册 handler
	for _, h := range lnet.Handlers {
		s.Register(h)
	}
	return s, nil
}
