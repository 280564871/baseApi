package jwt

import (
	"api/conf"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Info struct {
	Id int
	jwt.StandardClaims
}

func GenerateToken(id int) (string, error) {
	info := Info{
		id,
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 72).Unix()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, info)

	return token.SignedString(conf.AppCfg.Secret)
}

func CheckToken(tokenString string) (Info, error) {
	info := Info{}
	_, err := jwt.ParseWithClaims(tokenString, &info, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(conf.AppCfg.Secret), nil
	})
	if err != nil {
		return info, err
	}
	return info, nil
}
