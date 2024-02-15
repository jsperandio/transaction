package server

import (
	"log/slog"
	"os"

	fxsvc "github.com/jsperandio/transaction/app/fx/module/domain/service"
	fxprvdr "github.com/jsperandio/transaction/app/fx/module/provider"
	fxrest "github.com/jsperandio/transaction/app/fx/module/server/rest"
	fxrestecho "github.com/jsperandio/transaction/app/fx/module/server/rest/echo"
	"go.uber.org/fx"
)

func Start() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	app := fx.New(
		fx.Options(
			fxrestecho.Module(),
		),
		fxprvdr.Module(),
		fxsvc.Module(),
		fxrest.Module(),
		fx.Invoke(fxrestecho.ListenAndServe),
	)
	app.Run()
}
