package restfx

import (
	fxhndlr "github.com/jsperandio/transaction/app/fx/module/server/rest/handler"
	"go.uber.org/fx"
)

//	@title			transaction server
//	@version		1.0
//	@description	This is a implementation of transaction server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8081
// @BasePath	/
func Module() fx.Option {
	return fx.Options(
		fxhndlr.Module(),
	)
}
