package api

import (
	"github.com/Komsorn-sookdang/go-template/internal/pkg/engine/assessment"
	"github.com/Komsorn-sookdang/go-template/pkg/model/response_model"
)

type AssessmentService interface {
	GetAssessmentByID(assessmentID int64) (response_model.GetAssessmentByIDResponse, error)
}

type assessmentServiceImpl struct {
	assessmentEngine assessment.AssessmentEngine
}

func (s assessmentServiceImpl) GetAssessmentByID(assessmentID int64) (response_model.GetAssessmentByIDResponse, error) {
	d, err := s.assessmentEngine.FindAssessmentByID(assessmentID)
	if err != nil {
		return response_model.GetAssessmentByIDResponse{}, err
	}
	return response_model.GetAssessmentByIDResponse{
		AssessmentID: d.AssessmentID,
		Name:         d.Name,
	}, nil
}

func NewAssessmentService(assessmentEngine assessment.AssessmentEngine) AssessmentService {
	return &assessmentServiceImpl{
		assessmentEngine: assessmentEngine,
	}
}
