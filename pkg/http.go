package pkg

import "github.com/gin-gonic/gin"

type AppState struct {
	TraceID string
}

// ExtractAppState extracts the key "state" from the context and returns it if it exists, otherwise returns
// an empty state with a false boolean.
func ExtractAppState(c *gin.Context) (AppState, bool) {
	reqState, exists := c.Get("state")

	if !exists {
		return AppState{}, false
	}

	state, ok := reqState.(AppState)

	return state, ok
}

// AbortWithError aborts the current request and returns true if the error is not nil, otherwise returns false.
func AbortWithError(c *gin.Context, err error) bool {
	if err != nil {
		_ = c.Error(err)
		c.Abort()

		return true
	}

	return false
}
