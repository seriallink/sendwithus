package swu

type ErrorMessage struct {
	Code    int
	Message string
}

func (e *ErrorMessage) Error() string {
	return e.Message
}
