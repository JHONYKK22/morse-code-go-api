package dto

type Response struct {
	Data string `json:"data"`
}

type Request struct {
	Text string `json:"text" validate:"required"`
}

type ErrorResponse struct {
	Message string `json:"message" validate:"required"`
}
