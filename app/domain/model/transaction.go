package model

import "time"

type Transaction struct {
	ID              int
	AccountID       int
	OperationTypeID OperationType
	Amount          float64
	EventDate       time.Time
}
