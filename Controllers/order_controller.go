package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	models "order-api/Models"
	service "order-api/Service"
)

type orderController struct {
	os service.OrderServiceAPI
}

func ProvideOrderController(os service.OrderServiceAPI) *orderController {
	return &orderController{os: os}
}

func (oc orderController) CreateOrderController(c *gin.Context) {
	var (
		order  models.Order
		result gin.H
	)

	err := c.ShouldBindJSON(&order)

	if err != nil {
		c.JSON(http.StatusBadRequest, result)
		return
	}

	res := oc.os.CreateOrderService(c, order)
	c.JSON(200, res)
}

func (oc orderController) GetOrderController(c *gin.Context) {
	var order []models.Order

	res := oc.os.GetOrderService(c, order)

	c.JSON(200, res)
}

func (oc orderController) UpdateOrderController(c *gin.Context) {
	var (
		order  models.Order
		result gin.H
	)

	err := c.ShouldBindJSON(&order)

	if err != nil {
		c.JSON(http.StatusBadRequest, result)
		return
	}

	res := oc.os.UpdateOrderService(c, order)

	c.JSON(200, res)
}

func (oc orderController) DeleteOrderController(c *gin.Context) {
	res := oc.os.DeleteOrderService(c)
	fmt.Println(res)
	c.JSON(200, res)
}
