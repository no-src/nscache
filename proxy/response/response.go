package response

import (
	"github.com/no-src/nscache/proxy"
	"github.com/no-src/nsgo/jsonutil"
)

// Response the response of proxy
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

// NewResponse create a new response
func NewResponse(code int, message string, data any) *Response {
	bData, err := jsonutil.Marshal(data)
	if err != nil {
		code = proxy.StatusError
		message = err.Error()
		bData = nil
	}
	return &Response{
		Code:    code,
		Message: message,
		Data:    string(bData),
	}
}

// NewSuccessResponse create a new success response
func NewSuccessResponse(data any) *Response {
	return NewResponse(proxy.StatusSuccess, "ok", data)
}

// NewErrorResponse create a new error response
func NewErrorResponse(message string) *Response {
	return NewResponse(proxy.StatusError, message, "")
}

// NewNilErrorResponse create a new nil error response
func NewNilErrorResponse() *Response {
	return NewResponse(proxy.StatusNilError, "nil error", "")
}

// Bytes get the response bytes
func (r *Response) Bytes() []byte {
	if r == nil {
		return nil
	}
	result, _ := jsonutil.Marshal(r)
	return result
}
