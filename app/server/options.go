package server

import (
	"log/slog"

	"github.com/jsperandio/transaction/app/config"
)

type Options struct {
	LogLevel string `json:"loglevel"`
	Port     string `json:"port"`
}

func DefaultOptions() (*Options, error) {
	opt := &Options{}

	err := config.UnmarshalWithPath("app.server", opt)
	if err != nil {
		slog.Error("failed to unmarshal server options", err)
		return nil, err
	}

	return opt, nil
}
