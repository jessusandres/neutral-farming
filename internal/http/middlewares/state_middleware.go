package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"looker.com/neutral-farming/pkg"
)

func NewState() gin.HandlerFunc {
	return func(c *gin.Context) {
		newUuid, _ := uuid.NewUUID()
		reqUuid := newUuid.String()

		appState := pkg.AppState{
			TraceID: reqUuid,
		}

		c.Set("state", appState)

		c.Writer.Header().Set("X-Trace-Id", reqUuid)

		c.Next()
	}
}
