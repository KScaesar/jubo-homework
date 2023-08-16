package application

import (
	"context"

	"github.com/KScaesar/jubo-homework/backend/domain"
	"github.com/KScaesar/jubo-homework/backend/util"
)

type PatientService interface {
	QueryPatientList(ctx context.Context, dto *domain.QryPatientParam) (util.ListResponse[domain.PatientResponse], error)
}

func NewPatientUseCase(repo domain.PatientRepo) *PatientUseCase {
	return &PatientUseCase{repo: repo}
}

type PatientUseCase struct {
	repo domain.PatientRepo
}

func (uc *PatientUseCase) QueryPatientList(
	ctx context.Context, dto *domain.QryPatientParam,
) (
	util.ListResponse[domain.PatientResponse], error,
) {
	dto.PreProcess(true)
	return uc.repo.QueryPatientList(ctx, dto)
}
