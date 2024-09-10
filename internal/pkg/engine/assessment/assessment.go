package assessment

import (
	"github.com/Komsorn-sookdang/go-template/internal/pkg/dto"
	"github.com/Komsorn-sookdang/go-template/internal/pkg/repository"
)

type AssessmentEngine interface {
	FindAssessmentByID(assessmentID int64) (dto.AssessmentInfo, error)
}

type assessmentEngineImpl struct {
	assessmentRepo repository.AssessmentRepository
}

func (a *assessmentEngineImpl) FindAssessmentByID(assessmentID int64) (dto.AssessmentInfo, error) {
	assessmentEntity, err := a.assessmentRepo.FindAssessmentByID(assessmentID)
	if err != nil {
		return dto.AssessmentInfo{}, err
	}

	return dto.AssessmentInfo{
		AssessmentID: assessmentEntity.AssessmentID,
		Name:         assessmentEntity.Name,
	}, nil
}

func NewAssessmentEngine(assessmentRepo repository.AssessmentRepository) AssessmentEngine {
	return &assessmentEngineImpl{
		assessmentRepo: assessmentRepo,
	}
}
