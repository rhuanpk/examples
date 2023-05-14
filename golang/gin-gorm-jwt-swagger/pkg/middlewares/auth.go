package middlewares

import (
	"net/http"
	"strings"

	"backend/internal/consts"
	"backend/pkg/utils/errors"

	"github.com/gin-gonic/gin"
	"github.com/go-hl/jwt"
)

// Auth is the middleware that's validate if request have a token and if its valid.
func Auth(context *gin.Context) {

	token := strings.TrimPrefix(context.GetHeader("Token"), "Bearer ")
	if token == "" {
		context.JSON(
			http.StatusBadRequest,
			errors.HTTP{
				Code:  http.StatusBadRequest,
				Error: "request does not contain an access token",
			},
		)
		context.Abort()
		return
	}

	if isValid, err := jwt.IsValid(token, consts.JWTSECRETKEY); !isValid {
		context.JSON(
			http.StatusUnauthorized,
			errors.HTTP{
				Code:  http.StatusUnauthorized,
				Error: err.Error(),
			},
		)
		context.Abort()
		return
	}

	context.Next()

}
