package main

import (
	"fmt"
	_ "github.com/akhmettolegen/texert/docs"
	"github.com/akhmettolegen/texert/internal/api"
	"os"
)

// @title Texert API Swagger
// @version v0.0.1
// @description Texert API Swagger Documentation.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url https://www.texert.kz
// @contact.email support@texert.kz

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v1
// @query.collection.format multi

// @Security OAuth2Application

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl http://143.198.96.88:3000/oauth/sign-in

//@x-extension-openapi {"example": "value on a json format"}

func main() {
	//cmd.Execute()

	server, err := api.NewServer()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	server.Start()
}
