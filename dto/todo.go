package dto

type TodoResponse struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Response struct {
	Message    string      `json:"message"`
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}
