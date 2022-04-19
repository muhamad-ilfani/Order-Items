package router

import (
	"gin/controller"

	"github.com/gin-gonic/gin"
)

func StartServer(db *controller.OrderDatas) *gin.Engine {
	router := gin.Default()

	router.POST("/orders", db.CreateOrder)
	router.GET("/orders", db.GetOrder)
	router.PUT("/orders/:id", db.UpdateOrder)
	router.DELETE("/orders/:id", db.DeleteOrder)
	return router
}
