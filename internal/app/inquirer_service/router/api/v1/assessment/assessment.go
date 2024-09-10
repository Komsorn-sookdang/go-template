package assessment

import "github.com/gin-gonic/gin"

type AssessmentController struct {
}

func NewAssessmentController() *AssessmentController {
	return &AssessmentController{}
}

func (ac *AssessmentController) InitEndpoints(r gin.IRouter) {
	r.GET("", ac.GetAssessmentByID)
	// r.POST("", ac.PostAssessment)
}

func (ac *AssessmentController) GetAssessmentByID(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Get Assessment"})
}
