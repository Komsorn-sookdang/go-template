package health

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Module("HealthRouter",
	fx.Provide(NewHealthCheckController),
	fx.Provide(NewHealthCheckRouter),
)

type HealthCheckRouter struct {
	healthController *HealthCheckController
}

func (h *HealthCheckRouter) InitEndpoints(r gin.IRouter) {
	actuator := r.Group("/actuator")
	actuator.GET("/", h.healthController.HealthCheck)
}

func NewHealthCheckRouter(healthController *HealthCheckController) *HealthCheckRouter {
	return &HealthCheckRouter{healthController: healthController}
}
