package server

import (
	fxrest "github.com/jsperandio/transaction/app/fx/module/server/rest"
	fxrestecho "github.com/jsperandio/transaction/app/fx/module/server/rest/echo"
	"go.uber.org/fx"
)

func Start() {
	app := fx.New(
		fx.Options(
			fxrestecho.Module(),
		),
		fxrest.Module(),
		fx.Invoke(fxrestecho.ListenAndServe),
	)
	app.Run()
}
