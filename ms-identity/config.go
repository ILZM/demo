package main

import (
	"os"

	"github.com/supertokens/supertokens-golang/recipe/dashboard"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"
)

var SUPERTOKENS_CORE_URL = os.Getenv("SUPERTOKENS_CORE_URL")
var SUPERTOKENS_APP_NAME = os.Getenv("SUPERTOKENS_APP_NAME")
var SUPERTOKENS_API_URL = os.Getenv("SUPERTOKENS_API_URL")
var SUPERTOKENS_WEBSITE_URL = os.Getenv("SUPERTOKENS_WEBSITE_URL")

var SuperTokensConfig = supertokens.TypeInput{
	Supertokens: &supertokens.ConnectionInfo{
		ConnectionURI: SUPERTOKENS_CORE_URL,
	},
	AppInfo: supertokens.AppInfo{
		AppName:       SUPERTOKENS_APP_NAME,
		APIDomain:     SUPERTOKENS_API_URL,
		WebsiteDomain: SUPERTOKENS_WEBSITE_URL,
	},
	RecipeList: []supertokens.Recipe{
		emailpassword.Init(nil),
		session.Init(nil),
		dashboard.Init(nil),
	},
}
