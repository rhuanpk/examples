package auth

import "backend/internal/infra/api"

func init() {
	auth := api.V1.Group("/auth/")
	{
		auth.GET(
			"/",
			get,
		)
		auth.POST(
			"/",
			login,
		)
	}
}
