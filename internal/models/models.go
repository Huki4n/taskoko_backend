package models

// Response - стандартный формат ответа
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
