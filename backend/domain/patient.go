package domain

import (
	"context"
	"time"

	"github.com/KScaesar/jubo-homework/backend/util"
)

type PatientRepo interface {
	QueryPatientList(ctx context.Context, dto *DtoQryPatientParam) (util.DtoListResponse[DtoPatientResponse], error)
	QueryPatientById(ctx context.Context, id string) (DtoPatientResponse, error)
	LockPatientById(ctx context.Context, id string) (Patient, error)

	CreatePatient(ctx context.Context, patient *Patient) error
	UpdatePatient(ctx context.Context, patient *Patient) error
	DeletePatient(ctx context.Context, patient *Patient) error
}

type Patient struct {
	Id        string    `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
}
