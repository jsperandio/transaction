package request

type FindAccount struct {
	AccountID int64 `json:"-" param:"accountId" validate:"required,gt=0"`
}
