package common

//DefaultResponse default payload response
type DefaultResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//NewInternalServerErrorResponse default internal server error response
func NewInternalServerErrorResponse() DefaultResponse {
	return DefaultResponse{
		500,
		"Internal server error",
	}
}

//NewExpiredTokenErrorResponse default internal server error response
func NewExpiredTokenErrorResponse() DefaultResponse {
	return DefaultResponse{
		403,
		"Token has expired",
	}
}

//NewNotFoundResponse default not found error response
func NewNotFoundResponse() DefaultResponse {
	return DefaultResponse{
		404,
		"Not found",
	}
}

//NewBadRequestResponse default not found error response
func NewBadRequestResponse() DefaultResponse {
	return DefaultResponse{
		400,
		"Bad request",
	}
}

//NewConflictResponse default not found error response
func NewConflictResponse() DefaultResponse {
	return DefaultResponse{
		409,
		"Data has been modified",
	}
}

//NewForbiddenResponse default for Forbidden error response
func NewForbiddenResponse() DefaultResponse {
	return DefaultResponse{
		403,
		"Forbidden",
	}
}

//NewMissingHeaderResponse bad request caused by header missing
func NewMissingHeaderResponse() DefaultResponse {
	return DefaultResponse{
		403,
		"Missing header data",
	}
}
