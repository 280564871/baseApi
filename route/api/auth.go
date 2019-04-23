package api

import (
	"api/moudel"
	"github.com/gin-gonic/gin"
)

func Login(c gin.Context)  {
	var user moudel.User
	 if c.BindJSON(&user) == nil{
		 if user {

		 }
	 }
}