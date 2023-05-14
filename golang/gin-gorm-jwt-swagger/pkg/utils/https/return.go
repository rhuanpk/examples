package https

import (
	"backend/pkg/utils/errors"

	"github.com/gin-gonic/gin"
)

// ReturnJSONError is the standard error return of API.
func ReturnJSONError(codeReturn int, errOrString any, context *gin.Context) {
	var message string
	if str, isStr := errOrString.(string); isStr {
		message = str
	} else if err, isErr := errOrString.(error); isErr {
		message = err.Error()
	}
	context.JSON(
		codeReturn,
		&errors.HTTP{
			Code:  codeReturn,
			Error: message,
		},
	)
	context.Abort()
}
