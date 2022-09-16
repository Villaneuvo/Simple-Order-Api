package main

import (
	config "order-api/Config"
	connection "order-api/Connection"
	controllers "order-api/Controllers"
	repository "order-api/Repository"
	service "order-api/Service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load an config
	cfg := config.LoadConfig()

	// Call the connection db
	db := connection.DBInitialize(cfg.Database.Username, cfg.Database.Password, cfg.Database.Port, cfg.Database.Name, cfg.Database.Host)

	// InDB := &controllers.InDB{DB: db}
	orderRepository := repository.ProvideRepositoryOrder(*db)
	orderService := service.OrderServiceProvider(orderRepository)
	orderController := controllers.ProvideOrderController(orderService)

	router := gin.Default()
	router.POST("/order", orderController.CreateOrderController)
	router.GET("/orders", orderController.GetOrderController)
	router.PUT("/order", orderController.UpdateOrderController)
	router.DELETE("/orders/:id", orderController.DeleteOrderController)

	router.Run(cfg.ServerPort)
}
