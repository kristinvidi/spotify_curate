package constants

import "fmt"

var (
	ErrUnauthorized  error = fmt.Errorf("not authorized")
	ErrUnsuccessful  error = fmt.Errorf("unsuccessful")
	ErrMissingConfig error = fmt.Errorf("missing config")
)
