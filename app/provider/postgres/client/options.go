package client

import "github.com/jsperandio/transaction/app/config"

type Options struct {
	DatabaseURL  string `json:"databaseurl"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	DatabaseName string `json:"databasename"`
}

func DefaultOptions() (*Options, error) {
	opt := &Options{}

	err := config.UnmarshalWithPath("app.provider.postgres.client", opt)
	if err != nil {
		return nil, err
	}

	// &Options{
	// 	DatabaseURL:  "postgresql://pismo:pismo@localhost:5432/pismo?sslmode=disable",
	// 	Username:     "pismo",
	// 	Password:     "pismo",
	// 	DatabaseName: "pismo",
	// }

	return opt, nil
}
