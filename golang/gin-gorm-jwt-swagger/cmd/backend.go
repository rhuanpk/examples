package main

import (
	// Built-in imports.
	"log"

	// Project imports.
	"backend/internal/consts"
	"backend/internal/infra/api"

	// Load imports.
	_ "backend/internal/process"
)

// Swagger:
//
//	@title						Backend API
//	@version					1.0
//	@description				This is a sample backend API.
//	@termsOfService				http://swagger.io/terms/
//
//	@contact.name				Rhuan Patriky
//	@contact.url				https://linktr.ee/rhuanpk
//	@contact.email				support@rhuanpk.com
//
//	@license.name				GPL 3.0
//	@license.url				https://www.gnu.org/licenses/gpl-3.0.en.html
//
//	@host						localhost:9999
//	@BasePath					/api/v1
//
//	@securityDefinitions.basic	BasicAuth
//
//	@externalDocs.description	OpenAPI Documentation
//	@externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	log.Fatalln(api.Router.Run(consts.APIPORT))
}
