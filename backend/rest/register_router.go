package rest

import (
	"github.com/gin-gonic/gin"

	"github.com/KScaesar/jubo-homework/backend/configs"
)

func RegisterRouter(
	cfg *configs.ProjectConfig,
	patientH *PatientHandler,
) *gin.Engine {

	gin.SetMode(gin.DebugMode)
	// gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.NoRoute(WarnApiRouteNotExist)
	router.Use(GetOrSetCorrelationId)

	v1 := router.Group("/v1/api")

	v1.GET("/patients", patientH.QueryPatientList)

	return router
}
