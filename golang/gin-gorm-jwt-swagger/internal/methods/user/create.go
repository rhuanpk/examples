package user

import (
	"errors"
	"net/http"

	"backend/internal/models"
	"backend/internal/operations/database"
	errorsutils "backend/pkg/utils/errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Swagger:
//
//	@Summary		CREATE
//	@Description	Create a new user in database.
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			Token	header		string		true	"Bearer token."
//	@Param			JSON	body		models.User	true	"Json request."
//	@Success		201		{object}	models.User
//	@Failure		409		{object}	errors.HTTP
//	@Failure		422		{object}	errors.HTTP
//	@Failure		500		{object}	errors.HTTP
//	@Router			/user [post]
func create(context *gin.Context) {

	var user = &models.User{}

	if context.ShouldBindJSON(&user) != nil {
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

	if user.HashPassword() != nil {
		context.JSON(
			http.StatusInternalServerError,
			&errorsutils.HTTP{
				Code:  http.StatusInternalServerError,
				Error: "failed to hash password",
			},
		)
		context.Abort()
		return
	}

	if err := database.Instance.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			context.JSON(
				http.StatusConflict,
				&errorsutils.HTTP{
					Code:  http.StatusConflict,
					Error: "user already existing",
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
		context.Abort()
		return
	}

	user.Password = nil
	context.JSON(http.StatusCreated, user)

}
