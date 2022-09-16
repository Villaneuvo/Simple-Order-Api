package service

import (
	"fmt"
	models "order-api/Models"
	repository "order-api/Repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ServiceOrderImplementation struct {
	repo repository.OrderRepositoryAPI
}

type OrderServiceAPI interface {
	CreateOrderService(c *gin.Context, order models.Order) gin.H
	GetOrderService(c *gin.Context, order []models.Order) gin.H
	DeleteOrderService(c *gin.Context) gin.H
	UpdateOrderService(c *gin.Context, order models.Order) gin.H
}

func (soi ServiceOrderImplementation) CreateOrderService(c *gin.Context, order models.Order) gin.H {
	res, err := soi.repo.CreateOrderRepository(order)

	if err != nil {
		panic(err)
	}

	result := gin.H{
		"result": res,
	}

	return result
}

func (soi ServiceOrderImplementation) GetOrderService(c *gin.Context, order []models.Order) gin.H {
	res, err := soi.repo.GetOrderRepository()
	if err != nil {
		panic(err)
	}

	result := gin.H{
		"result": res,
	}

	return result
}

func (soi ServiceOrderImplementation) DeleteOrderService(c *gin.Context) gin.H {
	var order models.Order

	id := c.Param("id")

	idn, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	order, err = soi.repo.FindOrderById(idn)
	if err != nil {
		panic(err)
	}

	_, err = soi.repo.DeleteOrderRepository(order)

	if err != nil {
		panic(err)
	}

	result := gin.H{
		"result": "Data deleted successfully",
		"count":  1,
	}

	return result

}

func (soi ServiceOrderImplementation) UpdateOrderService(c *gin.Context, order models.Order) gin.H {
	var (
		updateOrder models.Order
		result      gin.H
	)

	order, err := soi.repo.FindOrderById(order.ID)
	if err != nil {
		panic(err)
	}

	updateOrder = order
	res, err := soi.repo.UpdateOrderRepository(updateOrder)
	fmt.Println(res)
	if err != nil {
		result = gin.H{
			"result": "Error updating data",
		}
	} else {
		result = gin.H{
			"result": "Update data successfully",
		}
	}

	return result

}

func OrderServiceProvider(repo repository.OrderRepositoryAPI) *ServiceOrderImplementation {
	return &ServiceOrderImplementation{repo: repo}
}
