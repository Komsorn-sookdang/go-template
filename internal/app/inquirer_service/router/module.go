package router

import (
	"github.com/Komsorn-sookdang/go-template/internal/app/inquirer_service/router/api"
	"github.com/Komsorn-sookdang/go-template/internal/app/inquirer_service/router/health"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Module("BaseRouter",
	health.Module,
	api.Module,
)

func SetupBaseGinRouter(healthCheckRouter *health.HealthCheckRouter, apiRouter *api.ApiRouter) *gin.Engine {
	r := gin.New()

	healthCheckRouter.InitEndpoints(r)
	apiRouter.InitEndpoints(r)
	return r
}
