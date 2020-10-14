package dnas

// errorResponse represents an error from DNA Spaces
type errorResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// Err implements the error interface so we can have constant errors.
type Err string

func (e Err) Error() string {
	return string(e)
}

// Error Constants
// Cisco documents these as the only error responses they will emit.
const (
	ErrBadRequest    = Err("dnas: bad request")
	ErrUnauthorized  = Err("dnas: unauthorized request")
	ErrForbidden     = Err("dnas: forbidden")
	ErrInternalError = Err("dnas: internal error")
	ErrUnknown       = Err("dnas: unexpected error occurred")
)
