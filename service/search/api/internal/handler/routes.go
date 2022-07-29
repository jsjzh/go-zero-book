// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"book/service/search/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/search/do",
				Handler: searchHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/search/ping",
				Handler: pingHandler(serverCtx),
			},
		},
	)
}
