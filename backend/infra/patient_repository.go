package infra

import (
	"context"

	"gorm.io/gorm/clause"

	"github.com/KScaesar/jubo-homework/backend/domain"
	"github.com/KScaesar/jubo-homework/backend/util"
	"github.com/KScaesar/jubo-homework/backend/util/database"
	"github.com/KScaesar/jubo-homework/backend/util/errors"
)

const (
	PatientTableName = "patients"
)

func NewPatientRepository(db *database.WrapperGorm) *PatientRepository {
	return &PatientRepository{db: db}
}

type PatientRepository struct {
	db *database.WrapperGorm
}

func (repo *PatientRepository) QueryPatientList(
	ctx context.Context,
	dto *domain.QryPatientParam,
) (util.ListResponse[domain.PatientResponse], error) {
	db := repo.db.ChooseProcessor(ctx)
	db = db.Table(PatientTableName)
	return database.GormQueryListWithPagination[domain.PatientResponse](
		db,
		dto.FilterPatientParam,
		dto.PageParam,
		dto.SortParam,
	)
}

func (repo *PatientRepository) QueryPatientById(ctx context.Context, id string) (patient domain.PatientResponse, err error) {
	db := repo.db.ChooseProcessor(ctx)

	err = db.Table(PatientTableName).Where("id = ?", id).First(&patient).Error
	if err != nil {
		err = errors.WrapWithMessage(database.GormError(err), "patient id = %v", id)
		return
	}
	return
}

func (repo *PatientRepository) LockPatientById(ctx context.Context, id string) (patient domain.Patient, err error) {
	db := repo.db.ChooseProcessor(ctx)

	err = db.Table(PatientTableName).Where("id = ?", id).Clauses(clause.Locking{Strength: "UPDATE"}).First(&patient).Error
	if err != nil {
		err = errors.WrapWithMessage(database.GormError(err), "id = %v", id)
		return
	}
	return
}

func (repo *PatientRepository) CreatePatient(ctx context.Context, patient *domain.Patient) error {
	db := repo.db.ChooseProcessor(ctx)

	err := db.Table(PatientTableName).Create(patient).Error
	if err != nil {
		return database.GormError(err)
	}
	return nil
}

func (repo *PatientRepository) UpdatePatient(ctx context.Context, patient *domain.Patient) error {
	db := repo.db.ChooseProcessor(ctx)

	err := db.Table(PatientTableName).Where("id = ?", patient.Id).Save(patient).Error
	if err != nil {
		return database.GormError(err)
	}
	return nil
}

func (repo *PatientRepository) DeletePatient(ctx context.Context, patient *domain.Patient) error {
	db := repo.db.ChooseProcessor(ctx)

	err := db.Table(PatientTableName).Where("id = ?", patient.Id).Delete(patient).Error
	if err != nil {
		return database.GormError(err)
	}
	return nil
}
