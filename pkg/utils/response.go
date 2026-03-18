package utils

type APIResponse struct {
	Success bool `json:"success"`
	Data    any  `json:"data,omitempty"`
	Error   any  `json:"error,omitempty"`
}

func SuccessResponse(data any) *APIResponse {
	return &APIResponse{
		Success: true,
		Data:    data,
	}
}

func ErrorResponse(code int, err string) *APIResponse {
	return &APIResponse{
		Success: false,
		Error:   err,
	}
}
