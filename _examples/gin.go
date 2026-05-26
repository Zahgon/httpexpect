package examples

import (
	"github.com/gin-gonic/gin"
)

// Rid of debug output
func init() {
	gin.SetMode(gin.TestMode)
}

// GinHandler Create add /example route to gin engine
func GinHandler(r *gin.Engine) *gin.Engine {
	_ = "STUB: not implemented"
	// Add route to the gin engine
	return nil
}

// return gin engine with newly added route
