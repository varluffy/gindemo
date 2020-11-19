package ginwrap

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/varluffy/gindemo/pkg/errors"
	"github.com/varluffy/gindemo/pkg/logger"
	"net/http"
)

const (
	prefix     = "gindemo"
	UserIDKey  = prefix + "/user-id"
	ResBodyKey = prefix + "/res-body"
)

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetUserID(c *gin.Context) string {
	return c.GetString(UserIDKey)
}

func SetUserID(c *gin.Context, userID string) {
	c.Set(UserIDKey, userID)
}

func ParseJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		return errors.Wrap400Response(err, fmt.Sprintf("解析参数错误 - %s", err.Error()))
	}
	return nil
}

func ParseForm(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindWith(obj, binding.Form); err != nil {
		return errors.Wrap400Response(err, fmt.Sprintf("解析参数错误- %s", err.Error()))
	}
	return nil
}

func ResJSON(c *gin.Context, status int, v interface{}) {
	buf, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	c.Set(ResBodyKey, buf)
	c.Data(http.StatusOK, "application/json; charset=utf-8", buf)
	c.Abort()
}

func ResSuccess(c *gin.Context, v interface{}) {
	res := &Result{
		Code:    0,
		Message: "success",
		Data:    v,
	}
	if v == nil {
		res.Data = map[string]string{}
	}
	ResJSON(c, http.StatusOK, res)
}

func ResOK(c *gin.Context) {
	ResSuccess(c, &Result{
		Code:    0,
		Message: "success",
		Data:    map[string]string{},
	})
}

func ResError(c *gin.Context, err error, status ...int) {
	ctx := c.Request.Context()
	var res *errors.ResponseError
	if err != nil {
		if e, ok := err.(*errors.ResponseError); ok {
			res = e
		} else {
			res = errors.UnWrapResponse(errors.Wrap500Response(err))
		}
	} else {
		res = errors.UnWrapResponse(errors.ErrInternal)
	}

	if len(status) > 0 {
		res.StatusCode = status[0]
	}
	result := &Result{
		Code:    res.Code,
		Message: res.Message,
		Data:    map[string]string{},
	}

	if err := res.ERR; err != nil {
		logger.StartSpan(ctx, logger.SetSpanTitle("res error")).Errorf(err.Error())
	} else {
		logger.StartSpan(ctx, logger.SetSpanTitle("res warn")).Warnf("%+v", result)
	}

	ResJSON(c, res.StatusCode, result)
}
