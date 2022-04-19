package router

import (
	"gin/controller"

	"github.com/gin-gonic/gin"
)

func StartServer(db *controller.OrderDatas) *gin.Engine {
	router := gin.Default()

	router.POST("/user", db.CreateOrder)
	router.GET("/user", db.GetOrder)
	router.PUT("/user/:id", db.UpdateOrder)
	router.DELETE("/user/:id", db.DeleteOrder)
	return router
}
