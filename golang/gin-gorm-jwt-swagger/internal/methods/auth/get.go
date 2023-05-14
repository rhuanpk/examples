package auth

import (
	"net/http"
	"strings"

	"backend/internal/consts"
	"backend/pkg/utils/errors"

	"github.com/gin-gonic/gin"
	"github.com/go-hl/jwt"
)

// Swagger:
//
//	@Summary		ID
//	@Description	Get user ID inside the token (JWT).
//	@Tags			auth
//	@Param			Token	header		string	true	"Bearer token."
//	@Success		200		{string}	string	"OK"
//	@Failure		400		{object}	errors.HTTP
//	@Failure		500		{object}	errors.HTTP
//	@Router			/auth [get]
func get(context *gin.Context) {

	token := strings.TrimPrefix(context.GetHeader("Token"), "Bearer ")
	if token == "" {
		context.JSON(
			http.StatusBadRequest,
			&errors.HTTP{
				Code:  http.StatusBadRequest,
				Error: "request does not contain an access token",
			},
		)
		context.Abort()
		return
	}

	ID, err := jwt.GetUserID(token, consts.JWTSECRETKEY)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			&errors.HTTP{
				Code:  http.StatusInternalServerError,
				Error: err.Error(),
			},
		)
		context.Abort()
		return
	}

	context.String(http.StatusOK, "User ID: %d", ID)

}
