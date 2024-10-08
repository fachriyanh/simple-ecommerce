package entity

import "errors"

var (
	ErrorNotFound     = errors.New("record not found")
	ErrorInternal     = errors.New("internal server error")
	ErrorNotAvailable = errors.New("error not available")
	ErrorConflict     = errors.New("error SKU already created")
	ErrBadRequest     = errors.New("error bad request")
)
