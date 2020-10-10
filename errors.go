package dnas

// Err implements the error interface so we can have constant errors.
type Err string

func (e Err) Error() string {
	return string(e)
}

// Error Constants
// Cisco documents these as the only error responses they will emit.
const (
	ErrBadRequest    = Err("bad request")
	ErrUnauthorized  = Err("unauthorized request")
	ErrForbidden     = Err("forbidden")
	ErrInternalError = Err("internal error")
	ErrUnknown       = Err("unexpected error occurred")
)
