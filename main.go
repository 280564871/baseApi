package main

import (
	"api/conf"
	"api/pkg"
	"api/route"
)

func main(){
	conf.InitConfig()
	pkg.InitDb()
	route.InitRoute()
	defer pkg.CloseDb()
}
