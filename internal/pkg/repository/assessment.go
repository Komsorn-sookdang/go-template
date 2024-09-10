package repository

import "github.com/Komsorn-sookdang/go-template/internal/pkg/entity"

type AssessmentRepository interface {
	FindAssessmentByID(assessmentID int64) (entity.Assessment, error)
}

type assessmentRepositoryImpl struct {
}

func (assessmentRepositoryImpl) FindAssessmentByID(assessmentID int64) (entity.Assessment, error) {
	return entity.Assessment{
		AssessmentID: assessmentID,
		Name:         "Aroi Sad",
	}, nil
}

func NewAssessmentRepository() AssessmentRepository {
	return &assessmentRepositoryImpl{}
}
