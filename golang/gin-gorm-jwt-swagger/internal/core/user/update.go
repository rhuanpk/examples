package user

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"

	"backend/internal/infra/database"
	"backend/internal/models"
	errorsutils "backend/pkg/utils/errors"

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
//	@Failure		400		{object}	errors.HTTP
//	@Failure		422		{object}	errors.HTTP
//	@Failure		500		{object}	errors.HTTP
//	@Router			/user/:id [patch]
func update(context *gin.Context) {

	var (
		user    = &models.User{}
		payload = make(map[string]any)
	)

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

	_, err := strconv.Atoi(ID)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			&errorsutils.HTTP{
				Code:  http.StatusBadRequest,
				Error: "wrong ID type",
			},
		)
		context.Abort()
		return
	}

	if json.NewDecoder(context.Request.Body).Decode(&user) != nil {
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

	wrapped := reflect.ValueOf(*user)
	for index := 0; index < wrapped.NumField(); index++ {
		fieldIndex := wrapped.Field(index)
		if fieldIndex.Kind() == reflect.Ptr {
			if fieldIndex.IsNil() {
				continue
			} else {
				fieldIndex = fieldIndex.Elem()
			}
		} else {
			if fieldIndex.IsZero() {
				continue
			}
		}
		payload[wrapped.Type().Field(index).Name] = fieldIndex.Interface()
	}

	if database.Instance.Model(&models.User{}).Where("id = ?", ID).Updates(payload).Error != nil {
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

	context.String(http.StatusNoContent, "successfully updated")

}
