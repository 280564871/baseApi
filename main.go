package main

import (
	"api/conf"
	"api/moudel"
	"api/route"
)

func main() {
	conf.InitConfig()
	moudel.InitDb()
	route.InitRoute()
}
