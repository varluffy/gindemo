package errors

import (
	"github.com/pkg/errors"
)

// 定义别名
var (
	New          = errors.New
	Wrap         = errors.Wrap
	Wrapf        = errors.Wrapf
	WithStack    = errors.WithStack
	WithMessage  = errors.WithMessage
	WithMessagef = errors.WithMessagef
)

var (
	ErrBadRequest        = New400Response("请求错误")
	ErrInternal          = New500Response("系统内部错误")
	ErrorMethodNotAllow  = NewResponse(405, "该请求方法不被允许")
	ErrorTooManyRequests = NewResponse(200, "请求过于频繁")
	ErrNotFound          = NewResponse(404, "资源不存在", 404)
)
