package usecase

import "errors"

var (
	ErrEmpty          = errors.New("empty cache")
	ErrChoosing       = errors.New("choose program")
	ErrOnlyOneProgram = errors.New("choose only 1 program")
	ErrLowInitPay     = errors.New("the initial payment should be more")
)
