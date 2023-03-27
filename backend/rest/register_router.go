package rest

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(
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
