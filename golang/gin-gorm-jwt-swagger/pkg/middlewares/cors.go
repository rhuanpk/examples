package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORS is the middleware that's validate if request have a token and if its valid.
func CORS(context *gin.Context) {
	context.Header("Access-Control-Allow-Origin", "*")
	context.Header("Access-Control-Allow-Methods", "*")
	context.Header("Access-Control-Allow-Headers", "*")
	if context.Request.Method == http.MethodOptions {
		context.Status(http.StatusContinue)
	}
	context.Next()
}
