package errors

// errors are every struct that implelments the error interface
// which means you have to impelement the Error function for it

// not found record in the database
type NotFoundError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	// ID       string
	// Location string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

func NewNotFoundError(message string) *NotFoundError {
	err := &NotFoundError{
		Code:    404,
		Message: message,
	}

	return err
}
