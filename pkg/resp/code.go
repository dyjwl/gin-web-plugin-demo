package resp

const (
	Success       = 0
	ErrBadRequest = 10400
	// ErrInvalidAuthHeader - 401: Invalid authorization header.
	ErrInvalidAuthHeader = 10401
	ErrSignatureInvalid  = 10402
	ErrPermissionDenied  = 10405
	ErrPageNotFound      = 10404
	ErrBind              = 10406
	ErrValidation        = 10407
	ErrInternalServer    = 10500
)
