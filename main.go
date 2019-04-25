package main

import (
	"api/conf"
	"api/models"
	"api/route"
)

func main() {
	conf.InitConfig()
	models.InitDb()
	route.InitRoute()
}
