package api

import (
	"backend/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

// Router is the singleton instance of api router.
var Router = gin.Default()

// V1 is the singleton instance of v1 route api group.
var V1 = Router.Group(
	"/api/v1",
	middlewares.CORS,
)

func init() {
	V1.OPTIONS("/*any")
}
