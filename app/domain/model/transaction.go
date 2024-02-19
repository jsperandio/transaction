package model

import "time"

type Transaction struct {
	ID              int64
	AccountID       int64
	OperationTypeID OperationType
	Amount          float64
	EventDate       time.Time
}
