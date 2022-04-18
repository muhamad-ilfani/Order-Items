package main

import (
	"gin/config"
	"gin/controller"
	"gin/router"
)

func main() {
	var PORT = ":8080"
	db := config.InitDB()

	inDB := &controller.OrderDatas{DB: db}

	router.StartServer(inDB).Run(PORT)
}
