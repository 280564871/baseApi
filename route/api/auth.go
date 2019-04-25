package api

import (
	"api/conf"
	"github.com/gin-gonic/gin"
	"github.com/medivhzhan/weapp"
	"log"
	"net/http"
)

func LoginByWx(c *gin.Context) {
	var request struct {
		Code     string `json:"code"`
		UserInfo struct {
			RawData       string `json:"rawData"`
			Signature     string `json:"signature"`
			EncryptedData string `json:"encryptedData"`
			Iv            string `json:"iv"`
		} `json:"userInfo"`
	}
	err := c.BindJSON(&request)
	user := request.UserInfo
	//code2session
	res, err := weapp.Login(conf.WxCfg.AppId, conf.WxCfg.AppSecret, request.Code)
	if err != nil {
		c.JSON(http.StatusBadRequest, Res{Msg: err.Error()})
		return
	}
	//验签用户信息
	userInfo, err := weapp.DecryptUserInfo(user.RawData, user.EncryptedData, user.Signature, user.Iv, res.SessionKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, Res{Msg: err.Error()})
		return
	}
	log.Println(userInfo)
	if err != nil {
		log.Println(err)
	}
	ReturnSuccess(c, "")
	return
}
