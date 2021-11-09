package errors

type RestError struct {
	Message	string	`json:"message"`
	Status	int		`json:"code"`
}

func NewRestError(m string, c int) *RestError {
	return &RestError{
		Message: m,
		Status: c,
	}
}