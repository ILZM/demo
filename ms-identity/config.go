package main

import (
	"log"
	"os"

	"github.com/supertokens/supertokens-golang/recipe/dashboard"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func GetSuperTokensConfig() supertokens.TypeInput {
	var SUPERTOKENS_CORE_URL = os.Getenv("SUPERTOKENS_CORE_URL")
	var SUPERTOKENS_APP_NAME = os.Getenv("SUPERTOKENS_APP_NAME")
	var SUPERTOKENS_API_URL = os.Getenv("SUPERTOKENS_API_URL")
	var SUPERTOKENS_WEBSITE_URL = os.Getenv("SUPERTOKENS_WEBSITE_URL")
	var SUPERTOKENS_API_BASE_PATH = os.Getenv("SUPERTOKENS_API_BASE_PATH")
	var SUPERTOKENS_API_GATEWAY_PATH = os.Getenv("SUPERTOKENS_API_GATEWAY_PATH")

	log.Println("SUPERTOKENS_CORE_URL: " + SUPERTOKENS_CORE_URL)
	log.Println("SUPERTOKENS_APP_NAME: " + SUPERTOKENS_APP_NAME)
	log.Println("SUPERTOKENS_API_URL: " + SUPERTOKENS_API_URL)
	log.Println("SUPERTOKENS_WEBSITE_URL: " + SUPERTOKENS_WEBSITE_URL)
	log.Println("SUPERTOKENS_API_BASE_PATH: " + SUPERTOKENS_API_BASE_PATH)
	log.Println("SUPERTOKENS_API_GATEWAY_PATH: " + SUPERTOKENS_API_GATEWAY_PATH)

	return supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			ConnectionURI: SUPERTOKENS_CORE_URL,
		},
		AppInfo: supertokens.AppInfo{
			AppName:        SUPERTOKENS_APP_NAME,
			APIDomain:      SUPERTOKENS_API_URL,
			WebsiteDomain:  SUPERTOKENS_WEBSITE_URL,
			APIBasePath:    &SUPERTOKENS_API_BASE_PATH,
			APIGatewayPath: &SUPERTOKENS_API_GATEWAY_PATH,
		},
		RecipeList: []supertokens.Recipe{
			emailpassword.Init(nil),
			session.Init(nil),
			dashboard.Init(nil),
		},
	}
}
