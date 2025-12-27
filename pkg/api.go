package pkg

import "github.com/gin-gonic/gin"

type AppState struct {
	Uuid string
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
