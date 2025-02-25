package rest_err

import "net/http"

type RestErr struct {
	Message string  `json:"message"`
	Err     string  `json:"error"`
	Code    int     `json:"code"`
	Causes  []Cause `json:"causes"`
}

type Cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewRestErr(message, err string, code int, causes []Cause) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  []Cause{},
	}
}
