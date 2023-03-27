//go:build wireinject

package ioc

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/KScaesar/jubo-homework/backend/configs"
	"github.com/KScaesar/jubo-homework/backend/rest"
	"github.com/KScaesar/jubo-homework/backend/util/database"
)

//go:generate wire gen

func NewHttpServerV1(cfg *configs.ProjectConfig) (*gin.Engine, error) {
	panic(wire.Build(
		infraDependency,
		appV1,

		rest.NewPatientHandler,
		rest.NewOrderHandler,

		rest.RegisterRouter,
	))
}

func NewHttpServerV2(cfg *configs.ProjectConfig) (*gin.Engine, error) {
	panic(wire.Build(
		infraDependency,

		NewAppV2,
		wire.FieldsOf(new(*AppV2), "PatientService"),
		wire.FieldsOf(new(*AppV2), "OrderService"),

		rest.NewPatientHandler,
		rest.NewOrderHandler,

		rest.RegisterRouter,
	))
}

func NewAppV2(cfg *configs.ProjectConfig, db *database.WrapperGorm) (*AppV2, error) {
	panic(wire.Build(
		appV2,
	))
}
