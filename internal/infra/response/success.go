package response

import "net/http"

func newSuccess(httpStatusCode int, msg string, opts ...OptResponse) Response {
	var resp = Response{
		HttpStatus: httpStatusCode,
		Message:    msg,
		Success:    true,
	}

	for _, opt := range opts {
		opt(&resp)
	}

	return resp
}

func NewSuccessCreated(msg string, opts ...OptResponse) Response {
	return newSuccess(http.StatusCreated, msg, opts...)
}

func NewSuccesOk(msg string, opts ...OptResponse) Response {
	return newSuccess(http.StatusOK, msg, opts...)
}
