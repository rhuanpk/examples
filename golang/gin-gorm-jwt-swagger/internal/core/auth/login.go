package auth

import (
	"errors"
	"net/http"
	"time"

	"backend/internal/consts"
	"backend/internal/infra/database"
	"backend/internal/models"
	errorsutils "backend/pkg/utils/errors"

	"github.com/gin-gonic/gin"
	"github.com/go-hl/jwt"
	"gorm.io/gorm"
)

// Swagger:
//
//	@Summary		LOGIN
//	@Description	Log-in and get a authentication token (JWT).
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			JSON	body		models.Auth	true	"Json request."
//	@Success		200		{object}	models.Token
//	@Failure		401		{object}	errors.HTTP
//	@Failure		404		{object}	errors.HTTP
//	@Failure		422		{object}	errors.HTTP
//	@Failure		500		{object}	errors.HTTP
//	@Router			/auth [post]
func login(context *gin.Context) {

	var (
		request = &models.Auth{}
		user    = &models.User{}
	)

	if context.ShouldBindJSON(&request) != nil {
		context.JSON(
			http.StatusUnprocessableEntity,
			&errorsutils.HTTP{
				Code:  http.StatusUnprocessableEntity,
				Error: "malformed JSON",
			},
		)
		context.Abort()
		return
	}

	if err := database.Instance.Where(
		"username = ? OR email = ?", request.Username, request.Email,
	).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(
				http.StatusNotFound,
				&errorsutils.HTTP{
					Code:  http.StatusNotFound,
					Error: "user not found",
				},
			)
		} else {
			context.JSON(
				http.StatusInternalServerError,
				&errorsutils.HTTP{
					Code:  http.StatusInternalServerError,
					Error: err.Error(),
				},
			)
		}
		return
	}

	if !user.IsValidPassword(*request.Password) {
		context.JSON(
			http.StatusUnauthorized,
			&errorsutils.HTTP{
				Code:  http.StatusUnauthorized,
				Error: "invalid password",
			},
		)
		context.Abort()
		return
	}

	token, err := jwt.Generate(
		time.Now().Add(time.Second*5),
		*user.ID,
		consts.JWTSECRETKEY,
	)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			&errorsutils.HTTP{
				Code:  http.StatusInternalServerError,
				Error: err.Error(),
			},
		)
		context.Abort()
		return
	}

	context.JSON(
		http.StatusOK,
		&models.Token{
			ID:    *user.ID,
			Token: token,
		},
	)

}
