package client

type Options struct {
	DatabaseURL  string `json:"databaseurl"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	DatabaseName string `json:"databasename"`
}

func DefaultOptions() *Options {
	return &Options{
		DatabaseURL:  "postgresql://pismo:pismo@localhost:5432/pismo?sslmode=disable",
		Username:     "pismo",
		Password:     "pismo",
		DatabaseName: "pismo",
	}
}
