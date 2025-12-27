package pkg

import "github.com/gin-gonic/gin"

type ApiError struct {
	Message string `json:"message"`
	Details any    `json:"payload,omitempty"`
	UUID    string `json:"uuid,omitempty"`
}

func (e *ApiError) ToResponse(c *gin.Context, statusCode int) {
	c.AbortWithStatusJSON(statusCode, gin.H{
		"error": ApiError{
			Message: e.Message,
			UUID:    e.UUID,
			Details: e.Details,
		},
	})
}
