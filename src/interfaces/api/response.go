package api

type ErrorResponse struct {
	Message string `json:"message"`
}

func GetErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Message: err.Error(),
	}
}
