package user

import (
	"net/http"

	"backend/internal/infra/database"
	"backend/internal/models"
	"backend/pkg/utils/errors"

	"github.com/gin-gonic/gin"
)

// Swagger:
//
//	@Summary		LIST
//	@Description	List all users in table.
//	@Tags			user
//	@Produce		json
//	@Param			Token	header		string	true	"Bearer token."
//	@Success		200		{array}		models.User
//	@Failure		500		{object}	errors.HTTP
//	@Router			/user [get]
func list(context *gin.Context) {

	var users = []*models.User{}

	if err := database.Instance.Omit(omittedColumn).Find(&users).Error; err != nil {
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

	context.JSON(http.StatusOK, users)

}
