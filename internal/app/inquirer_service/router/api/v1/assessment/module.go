package assessment

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type AssessmentRouter struct {
	assessmentController *AssessmentController
}

var Module = fx.Module("ApiV1AssessmentModule",
	fx.Provide(NewAssessmentController),
	fx.Provide(NewAssessmentRouter),
)

func NewAssessmentRouter(assessmentController *AssessmentController) *AssessmentRouter {
	return &AssessmentRouter{assessmentController: assessmentController}
}

func (ar AssessmentRouter) InitEndpoints(r gin.IRouter) {
	assessmentGroup := r.Group("/assessment")
	ar.assessmentController.InitEndpoints(assessmentGroup)
}
