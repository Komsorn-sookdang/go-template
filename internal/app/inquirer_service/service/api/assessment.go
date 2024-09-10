package api

import (
	"github.com/Komsorn-sookdang/go-template/internal/pkg/engine/assessment"
)

type AssessmentService interface {
	GetAssessmentByID(assessmentID int64) (response_model.GetAssessmentByID, error)
}

type assessmentServiceImpl struct {
	assessmentEngine assessment.AssessmentEngine
	// assessmentService
}

func (s AssessmentService) GetAssessmentByID(assessmentID int64) (response_model.GetAssessmentByID, error) {
	d, err := s.assessmentEngine.GetAssessmentByID(assessmentID)
	if err != nil {
		return response_model.GetAssessmentByID{}, err
	}
	return response_model.GetAssessmentByID{}
}

func NewAssessmentService(assessmentEngine assessment.AssessmentEngine) AssessmentService {
	return &assessmentServiceImpl{
		assessmentEngine: assessmentEngine,
	}
}

// https://github.com/Komsorn-sookdang/go-template.git
