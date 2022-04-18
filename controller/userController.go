package controller

import (
	"gin/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderDatas struct {
	DB *gorm.DB
}

func (idb *OrderDatas) CreateOrder(ctx *gin.Context) {
	var newOrder model.Order

	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := idb.DB.Create(&newOrder).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"massage": "failed to create order data",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"order": newOrder,
	})
}

func (idb *OrderDatas) GetOrder(ctx *gin.Context) {
	var dataOrder []model.Order

	if err := idb.DB.Find(&dataOrder).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status": "Data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"order": dataOrder,
	})
}

/*
func (idb *CarDatas) UpdateCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	var Data model.Car

	if err := ctx.ShouldBindJSON(&Data); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := idb.DB.Model(&Data).Where("car_id=?", carID).Updates(Data).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"massage": "failed to update data",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"car": Data,
	})
}

func (idb *CarDatas) DeleteCar(ctx *gin.Context) {
	carID := ctx.Param("carID")

	var Data model.Car

	if err := ctx.ShouldBindJSON(&Data); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := idb.DB.Model(&Data).Where("car_id", carID).Delete(Data).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"massage": "failed to delete data",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car %v successfully deleted", carID),
	})
}
*/
