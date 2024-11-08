// Package usecase s
package usecase

import "errors"

// Errors.
var (
	ErrEmpty          = errors.New("empty cache")
	ErrChoosing       = errors.New("choose program")
	ErrOnlyOneProgram = errors.New("choose only 1 program")
	ErrLowInitPay     = errors.New("the initial payment should be more")

	ErrInvalidKeyType   = errors.New("invalid key type")
	ErrInvalidValueType = errors.New("invalid value type")
)
