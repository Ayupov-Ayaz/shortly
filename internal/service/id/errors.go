package id

import "errors"

var (
	ErrInvalidBase62String   = errors.New("invalid base 62 string")
	ErrValueToLargeForBase62 = errors.New("value to large for base 62")
)
