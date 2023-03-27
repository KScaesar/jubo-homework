package rest

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(
	patientH *PatientHandler,
	orderH *OrderHandler,
) *gin.Engine {

	gin.SetMode(gin.DebugMode)
	// gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.NoRoute(WarnApiRouteNotExist)
	router.Use(GetOrSetCorrelationId)

	v1 := router.Group("/v1/api")

	v1.GET("/patients", patientH.QueryPatientList)

	v1.GET("/orders", orderH.QueryOrderList)
	v1.GET("/orders/:order_id", orderH.QueryOrderById)
	v1.POST("/orders", orderH.CreateOrder)
	v1.PATCH("/orders/:order_id", orderH.UpdateOrderInfo)

	return router
}
