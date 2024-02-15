package request

type FindAccount struct {
	AccountID int `json:"-" param:"accountId" validate:"required,gt=0"`
}
