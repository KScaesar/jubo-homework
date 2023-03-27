package ioc

import (
	"github.com/google/wire"

	"github.com/KScaesar/jubo-homework/backend/application"
	"github.com/KScaesar/jubo-homework/backend/configs"
	"github.com/KScaesar/jubo-homework/backend/domain"
	"github.com/KScaesar/jubo-homework/backend/infra"
	"github.com/KScaesar/jubo-homework/backend/util/database"
)

// wire bug because generic type:
// https://github.com/google/wire/pull/360#issuecomment-1141376353

var infraDependency = wire.NewSet(
	wire.FieldsOf(new(*configs.ProjectConfig), "Pgsql"),

	database.NewGormPgsql,

	database.NewGormTxFactory,
	wire.Bind(new(database.TransactionFactory), new(*database.GormTxFactory)),
)

var appV1 = wire.NewSet(
	application.NewPatientUseCase,
	wire.Bind(new(application.PatientService), new(*application.PatientUseCase)),
	infra.NewPatientRepository,
	wire.Bind(new(domain.PatientRepo), new(*infra.PatientRepository)),
)

var appV2 = wire.NewSet(
	wire.Struct(new(AppV2), "*"),

	application.NewPatientUseCase,
	wire.Bind(new(application.PatientService), new(*application.PatientUseCase)),
	infra.NewPatientRepository,
	wire.Bind(new(domain.PatientRepo), new(*infra.PatientRepository)),
)

type AppV2 struct {
	PatientService application.PatientService
}
