package ginwrap

import (
	"github.com/varluffy/gindemo/pkg/errors"
)

var (
	ErrUserUnauthorized = errors.NewResponse(401, "用户未登录")
	ErrUserTokenExpired = errors.NewResponse(402, "用户token已过期")
)
