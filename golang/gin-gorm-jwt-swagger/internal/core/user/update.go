package user

import (
	"encoding/json"
	"net/http"

	"backend/internal/infra/database"
	"backend/internal/models"
	"backend/pkg/utils/errors"

	"github.com/gin-gonic/gin"
)

// Must send all field in JSON request (blank fields for "null" values).
//
// Swagger:
//
//	@Summary		UPDATE
//	@Description	Update already existing user in database.
//	@Description	OBS: send all fields in JSON request (blank fields for "null" values).
//	@Tags			user
//	@Accept			json
//	@Param			Token	header		string		true	"Bearer token."
//	@Param			JSON	body		models.User	true	"Json request."
//	@Param			id		path		int			true	"User ID."
//	@Success		204		{string}	string		"No Content"
//	@Failure		422		{object}	errors.HTTP
//	@Failure		500		{object}	errors.HTTP
//	@Router			/user/:id [patch]
func update(context *gin.Context) {

	var (
		payload = make(map[string]any)
		ID      = context.Param("id")
	)

	if json.NewDecoder(context.Request.Body).Decode(&payload) != nil {
		context.JSON(
			http.StatusUnprocessableEntity,
			&errors.HTTP{
				Code:  http.StatusUnprocessableEntity,
				Error: "malformed JSON",
			},
		)
		context.Abort()
		return
	}

	if err := database.Instance.Model(&models.User{}).Where("id = ?", ID).Updates(payload).Error; err != nil {
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

	context.String(http.StatusNoContent, "successfully updated")

}
