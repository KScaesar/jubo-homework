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

func NewHttpServer(cfg *configs.ProjectConfig) (*gin.Engine, error) {
	panic(wire.Build(
		infraDependency,
		rest.RegisterRouter,
		NewApp,

		rest.NewPatientHandler,
		wire.FieldsOf(new(*App), "PatientService"),
	))
}

func NewApp(cfg *configs.ProjectConfig, db *database.WrapperGorm) (*App, error) {
	panic(wire.Build(
		app,
	))
}
