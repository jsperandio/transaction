package server

import (
	"log/slog"
	"os"

	"github.com/jsperandio/transaction/app/config"
	fxsvc "github.com/jsperandio/transaction/app/fx/module/domain/service"
	fxprvdr "github.com/jsperandio/transaction/app/fx/module/provider"
	fxrest "github.com/jsperandio/transaction/app/fx/module/server/rest"
	fxrestecho "github.com/jsperandio/transaction/app/fx/module/server/rest/echo"
	"github.com/jsperandio/transaction/app/server"
	"go.uber.org/fx"
)

func logOverride() {
	opt, err := server.DefaultOptions()
	if err != nil {
		panic(err)
	}

	var ll slog.Level
	err = ll.UnmarshalText([]byte(opt.LogLevel))
	if err != nil {
		panic(err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: ll == slog.LevelDebug,
		Level:     ll,
	}))
	slog.SetDefault(logger)
}

func Start() {
	config.Load()
	logOverride()

	app := fx.New(
		fx.Options(
			fx.Provide(
				server.DefaultOptions,
			),
			fxrestecho.Module(),
		),
		fxprvdr.Module(),
		fxsvc.Module(),
		fxrest.Module(),
		fx.Invoke(fxrestecho.ListenAndServe),
	)
	app.Run()
}
