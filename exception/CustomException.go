package exception

func New(text string) error {
	return &errorString{text}
}

type Response struct {
	Message string `json:"message"`
}

func NewExceptionResponse(message string) *Response {
	return &Response{Message: message}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
