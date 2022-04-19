package controller

import (
	"fmt"
	"gin/model"
	"net/http"
	"strconv"

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

	if err := idb.DB.Create(&newOrder).Error; err != nil {
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

	if err := idb.DB.Preload("Item_user").Find(&dataOrder).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status": "Data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"order": dataOrder,
	})
}

func (idb *OrderDatas) UpdateOrder(ctx *gin.Context) {
	ID := ctx.Param("id")
	var newData model.Order
	var Data model.Order

	orderID, _ := strconv.Atoi(ID)

	if err := ctx.ShouldBindJSON(&newData); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	if err := idb.DB.Model(&Data.Item_user).Where("order_id=?", orderID).Updates(&newData.Item_user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"massage": "failed to update item",
			"error":   err.Error(),
		})
		return
	}
	if err := idb.DB.Model(&Data).Where("order_id=?", orderID).Updates(&newData).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"massage": "failed to update order",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"order": newData,
	})
}

func (idb *OrderDatas) DeleteOrder(ctx *gin.Context) {
	ID := ctx.Param("id")
	orderID, _ := strconv.Atoi(ID)
	var Data model.Order

	if err := ctx.ShouldBindJSON(&Data); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := idb.DB.Model(&Data).Where("order_id=?", orderID).Delete(Data).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"massage": "failed to delete data",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car %v successfully deleted", orderID),
	})
}
