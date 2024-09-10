package v1

import (
	"github.com/Komsorn-sookdang/go-template/internal/app/inquirer_service/router/api/v1/assessment"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type V1Router struct {
	assessmentRouter *assessment.AssessmentRouter
}

var Module = fx.Module("ApiV1Module",
	assessment.Module,
	fx.Provide(NewV1Router),
)

func NewV1Router(
	assessmentRouter *assessment.AssessmentRouter,
) *V1Router {
	return &V1Router{assessmentRouter: assessmentRouter}
}

func (v1Router V1Router) InitEndpoints(r gin.IRouter) {
	v1Group := r.Group("/v1")
	v1Router.assessmentRouter.InitEndpoints(v1Group)
}
