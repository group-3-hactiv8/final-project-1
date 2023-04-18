package utils

type HttpError struct {
	Message interface{} `json:"message"`
	Trace   interface{} `json:"stack_trace"`
}

func NewHttpError(message interface{}, trace interface{}) HttpError {
	return HttpError{
		Message: message,
		Trace:   trace,
	}
}

type ResponseData struct {
	Status  interface{}       `json:"status,omitempty"`
	Message map[string]string `json:"message,omitempty"`
	Details interface{}
}
