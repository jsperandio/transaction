package server

import (
	"log/slog"
	"os"

	"github.com/jsperandio/transaction/app/config"
	fxsvc "github.com/jsperandio/transaction/app/fx/module/domain/service"
	fxprvdr "github.com/jsperandio/transaction/app/fx/module/provider"
	fxrest "github.com/jsperandio/transaction/app/fx/module/server/rest"
	fxrestecho "github.com/jsperandio/transaction/app/fx/module/server/rest/echo"
	"go.uber.org/fx"
)

func logOverride() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))
	slog.SetDefault(logger)
}

func Start() {
	config.Load()
	logOverride()

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
