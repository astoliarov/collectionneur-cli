package interfaces

import "errors"

var (
	ErrBadApiResponse = errors.New("Bad API response")
	ErrApiNoInfo      = errors.New("No API info")
)
