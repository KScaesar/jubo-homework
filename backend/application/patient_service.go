package application

import (
	"context"

	"github.com/KScaesar/jubo-homework/backend/domain"
	"github.com/KScaesar/jubo-homework/backend/util"
)

type PatientService interface {
	QueryPatientList(ctx context.Context, dto *domain.DtoQryPatientParam) (util.DtoListResponse[domain.DtoPatientResponse], error)
}

func NewPatientUseCase(repo domain.PatientRepo) *PatientUseCase {
	return &PatientUseCase{repo: repo}
}

type PatientUseCase struct {
	repo domain.PatientRepo
}

func (uc *PatientUseCase) QueryPatientList(ctx context.Context, dto *domain.DtoQryPatientParam) (util.DtoListResponse[domain.DtoPatientResponse], error) {
	dto.DtoPageParam.SetDefault(1000)
	dto.DtoSortPatientParam.SetDefault()
	return uc.repo.QueryPatientList(ctx, dto)
}
