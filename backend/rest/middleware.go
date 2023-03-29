package rest

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/KScaesar/jubo-homework/backend/util"
	"github.com/KScaesar/jubo-homework/backend/util/errors"
)

func WarnApiRouteNotExist(c *gin.Context) {
	c.JSON(http.StatusNotFound, Response{
		Code:    errors.ErrNotFound.MyCode(),
		Message: "api route not exist, confirm your url or http method",
		Payload: struct{}{},
	})
}

func GetOrSetCorrelationId(c *gin.Context) {
	const CorrelationIdHeaderKey = "X-Correlation-Id"
	corId := c.Request.Header.Get(CorrelationIdHeaderKey)
	if corId == "" {
		corId = util.NewUlid()
	}
	c.Writer.Header().Add(CorrelationIdHeaderKey, corId)

	ctx := c.Request.Context()
	ctx1 := util.ContextWithCorrelationId(ctx, corId)
	c.Request = c.Request.WithContext(ctx1)

	c.Next()
}

func CorsMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Access-Control-Allow-Headers", "Authorization, Content-Type"},
		AllowCredentials: true,
		MaxAge:           8 * time.Hour,
		AllowWebSockets:  true,
	})
}
