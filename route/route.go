package route

import (
	"api/conf"
	"api/middlewear/jwt"
	"github.com/gin-gonic/gin"
)

//初始化路由
func InitRoute() {
	g := gin.New()
	g.Use(gin.Logger())

	g.Group("/v1", func(context *gin.Context) {
		g.Group("login")
		g.Use(checkToken)
	})
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
