package main

import (
	"Goim-server/router"
	"Goim-server/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	r := router.Router()
	r.Run()
}
