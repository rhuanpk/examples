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
//	@Summary		GET
//	@Description	Get a single user from ID.
//	@Tags			user
//	@Produce		json
//	@Param			Token	header		string	true	"Bearer token."
//	@Param			id		path		int		true	"User ID."
//	@Success		200		{object}	models.User
//	@Failure		400		{object}	errors.HTTP
//	@Failure		404		{object}	errors.HTTP
//	@Failure		500		{object}	errors.HTTP
//	@Router			/user/:id [get]
func get(context *gin.Context) {

	var user = &models.User{}

	ID := context.Param("id")
	if ID == "" {
		context.JSON(
			http.StatusBadRequest,
			&errorsutils.HTTP{
				Code:  http.StatusBadRequest,
				Error: "missing ID param",
			},
		)
		context.Abort()
		return
	}

	if err := database.Instance.Omit(omittedColumn).First(user, ID).Error; err != nil {
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
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, user)

}
