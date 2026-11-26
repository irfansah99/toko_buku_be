package errors

type AppError struct {
	Message string
	Code    string
	Status  int
	Errors  map[string]string
	Data    any
}

func (e *AppError) Error() string {
	return e.Message
}
