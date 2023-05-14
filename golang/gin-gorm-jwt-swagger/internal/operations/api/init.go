package api

import "github.com/gin-gonic/gin"

// Router is the singleton instance of api router.
var Router *gin.Engine

// V1 is the singleton instance of v1 route api group.
var V1 *gin.RouterGroup

func init() {
	Router = gin.Default()
	V1 = Router.Group("/api/v1")
}
