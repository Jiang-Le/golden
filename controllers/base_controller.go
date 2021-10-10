package controllers

import (
	cerror "github.com/jiel/golden/controllers/error"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidRequest = cerror.NewError(10001, "请求参数错误")

	ErrAlbumsQueryFail = cerror.NewError(11001, "图集查询失败")
	ErrAlbumCreateFail = cerror.NewError(11002, "图集创建失败")
)

type Response struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type BaseController struct {
}

func (c *BaseController) serveResponse(ctx *gin.Context, data interface{}, msg ...string) {
	rsp := &Response{
		Code: 0,
		Data: data,
	}
	if len(msg) > 0 {
		rsp.Msg = msg[0]
	}
	ctx.JSON(http.StatusOK, rsp)
}

func (c *BaseController) serveError(ctx *gin.Context, err cerror.IError, msg ...string) {
	errMsg := err.ErrorMsg()
	if len(msg) > 0 {
		errMsg = msg[0]
	}
	rsp := &Response{
		Code: err.ErrorCode(),
		Msg:  errMsg,
	}
	ctx.JSON(err.StatusCode(), rsp)
}
