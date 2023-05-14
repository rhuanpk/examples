package middlewares

import "github.com/gin-gonic/gin"

// CORS is the middleware that's validate if request have a token and if its valid.
func CORS(context *gin.Context) {
	context.Header("Access-Control-Allow-Origin", "*")
	context.Header("Access-Control-Allow-Methods", "*")
	context.Header("Access-Control-Allow-Headers", "*")
	context.Next()
}
