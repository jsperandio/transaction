package client

type Options struct {
	DatabaseURL  string `json:"databaseurl"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	DatabaseName string `json:"databasename"`
}

func DefaultOptions() *Options {
	return &Options{
		DatabaseURL:  "postgresql://pismo:pismo@transaction-pg:5432/pismo",
		Username:     "pismo",
		Password:     "pismo",
		DatabaseName: "pismo",
	}
}
