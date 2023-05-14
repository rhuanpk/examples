package user

import (
	"backend/internal/operations/api"
	"backend/pkg/middlewares"
)

const omittedColumn = "password"

func init() {
	user := api.V1.Group("/user", middlewares.Auth)
	{
		user.GET(
			"/",
			list,
		)
		user.GET(
			"/:id",
			get,
		)
		user.POST(
			"/",
			create,
		)
		user.PATCH(
			"/:id",
			update,
		)
	}
}
