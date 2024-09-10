package api

import (
	v1 "github.com/Komsorn-sookdang/go-template/internal/app/inquirer_service/router/api/v1"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type ApiRouter struct {
	v1Router *v1.V1Router
}

var Module = fx.Module("ApiController",
	v1.Module,
	fx.Provide(NewApiRouter),
)

func NewApiRouter(v1Router *v1.V1Router) *ApiRouter {
	return &ApiRouter{v1Router}
}

func (apiRouter *ApiRouter) InitEndpoints(r gin.IRouter) {
	apiGroup := r.Group("/api")
	apiRouter.v1Router.InitEndpoints(apiGroup)
}
