package assessment

import (
	"net/http"

	"github.com/Komsorn-sookdang/go-template/internal/app/inquirer_service/service/api"
	"github.com/Komsorn-sookdang/go-template/pkg/model/request_model/uri"
	"github.com/gin-gonic/gin"
)

type AssessmentController struct {
	assessmentService api.AssessmentService
}

func NewAssessmentController(assessmentService api.AssessmentService) *AssessmentController {
	return &AssessmentController{
		assessmentService: assessmentService,
	}
}

func (ac *AssessmentController) InitEndpoints(r gin.IRouter) {
	r.GET("/:assessmentID", ac.GetAssessmentByID)
	// r.POST("", ac.PostAssessment)
}

func (ac *AssessmentController) GetAssessmentByID(ctx *gin.Context) {
	// should bind the request to a struct
	var assessmentURI uri.AssessmentIDURI

	if err := ctx.ShouldBindUri(&assessmentURI); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	resp, err := ac.assessmentService.GetAssessmentByID(assessmentURI.AssessmentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Get Assessment", "data": resp, "shouldRemoveFromBFF": true})
}
