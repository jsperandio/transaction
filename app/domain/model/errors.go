package model

import "errors"

// ErrAccountAlreadyExists is returned when an account already exists with given documentnumber.
var ErrAccountAlreadyExists = errors.New("an account already exists with given documentnumber")
