package utils

type HttpSuccess[T any] struct {
	Message interface{} `json:"message"`
	Data    T           `json:"data"`
}

func NewHttpSuccess[T any](message interface{}, data T) HttpSuccess[T] {
	return HttpSuccess[T]{
		Message: message,
		Data:    data,
	}
}
