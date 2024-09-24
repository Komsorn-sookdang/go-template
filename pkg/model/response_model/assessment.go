package response_model

type GetAssessmentByIDResponse struct {
	AssessmentID int64  `json:"assessment_id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
}
