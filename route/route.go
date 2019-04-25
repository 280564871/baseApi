package route

import (
	"api/conf"
	"api/middlewear/jwt"
	"api/route/api"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

//初始化路由
func InitRoute() {
	g := gin.Default()
	f, _ := os.Create("log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	v1 := g.Group("/v1")
	{
		v1.POST("/wxLogin", api.LoginByWx)
		v1.Use(checkToken)
	}
	g.Run()
}

func checkToken(context *gin.Context) {
	var info jwt.Info
	var err error
	tokenString := context.GetHeader("token")
	info, err = jwt.CheckToken(tokenString)
	if err != nil {
		code := conf.AUTH_ERROR
		context.JSON(code, conf.ErrMsg[code])
	}
	context.Set("user", info)
}
