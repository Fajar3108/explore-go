package helpers

type ResponseHelper struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Meta    any    `json:"meta,omitempty"`
}

func NewResponseHelper(code int, message string, data any, meta any) *ResponseHelper {
	return &ResponseHelper{
		code,
		message,
		data,
		meta,
	}
}
