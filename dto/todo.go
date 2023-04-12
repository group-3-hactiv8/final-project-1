package dto

type TodoResponse struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
