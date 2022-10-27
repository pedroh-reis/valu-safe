package http

type HTTPDefaultResponseBody struct {
	Message string `json:"message"`
}

func NewHTTPDefaultResponseBody(message string) *HTTPDefaultResponseBody {
	return &HTTPDefaultResponseBody{
		Message: message,
	}
}
