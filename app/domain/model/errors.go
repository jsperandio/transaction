package model

import "errors"

var (
	// ErrAccountAlreadyExists is returned when an account already exists with given documentnumber.
	ErrAccountAlreadyExists = errors.New("an account already exists with given documentnumber")
	// ErrAccountNotFound is returned when an account is not found.
	ErrAccountNotFound = errors.New("account not found")
	// ErrInvalidTransaction is returned when an invalid transaction is provided.
	ErrInvalidTransaction = errors.New("invalid transaction")
	// ErrInvalidOperationType is returned when an invalid operation type is provided.
	ErrInvalidOperationType = errors.New("invalid operation type")
	// ErrNonAvaliableLimitForAccountTransaction return on limit is overpass
	ErrNonAvaliableLimitForAccountTransaction = errors.New("account limit is overrpass")
)
