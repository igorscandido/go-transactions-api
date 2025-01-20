package dto

type DefaultResponse struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message" example:"payment processed successfully"`
}
