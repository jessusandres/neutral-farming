package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"looker.com/neutral-farming/pkg"
)

func NotFound(c *gin.Context) {
	routePath := c.Request.URL.Path
	routeMethod := c.Request.Method

	message := fmt.Sprintf("Route [%s] %s not found", routeMethod, routePath)
	uuid := ""

	state, ok := pkg.ExtractAppState(c)

	if ok {
		uuid = state.Uuid
	}

	apiErr := pkg.ApiError{
		Message: message,
		UUID:    uuid,
	}

	apiErr.ToResponse(c, http.StatusNotFound)
}
