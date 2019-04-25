package api

import (
	"api/conf"
	"github.com/gin-gonic/gin"
)

type Res struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ReturnError(ctx *gin.Context, code int) {
	res := map[string]interface{}{
		"code": code,
		"msg":  conf.GetMsg(code),
		"data": "",
	}
	ctx.JSON(code, res)
}

func ReturnSuccess(ctx *gin.Context, data interface{}) {
	code := conf.SUCCESS
	res := map[string]interface{}{
		"code": code,
		"msg":  conf.GetMsg(code),
		"data": data,
	}
	ctx.JSON(code, res)
}
