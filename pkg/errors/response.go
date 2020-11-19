package errors

type ResponseError struct {
	Code       int
	Message    string
	StatusCode int
	ERR        error
}

func (r *ResponseError) Error() string {
	if r.ERR != nil {
		return r.ERR.Error()
	}
	return r.Message
}

func UnWrapResponse(err error) *ResponseError {
	if v, ok := err.(*ResponseError); ok {
		return v
	}
	return nil
}

func WrapResponse(err error, code int, msg string, status ...int) error {
	res := &ResponseError{
		Code:    code,
		Message: msg,
		ERR:     err,
	}
	if len(status) > 0 {
		res.StatusCode = status[0]
	}
	return res
}

func Wrap400Response(err error, msg ...string) error {
	m := "请求错误"
	if len(msg) > 0 {
		m = msg[0]
	}
	return WrapResponse(err, 400, m)
}

func Wrap500Response(err error, msg ...string) error {
	m := "系统内部错误"
	if len(msg) > 0 {
		m = msg[0]
	}
	return WrapResponse(err, 500, m)
}

func NewResponse(code int, msg string, status ...int) error {
	res := &ResponseError{
		Code:    code,
		Message: msg,
	}
	if len(status) > 0 {
		res.StatusCode = status[0]
	}
	return res
}

func New400Response(msg string) error {
	return NewResponse(400, msg)
}

func New500Response(msg string) error {
	return NewResponse(500, msg)
}
