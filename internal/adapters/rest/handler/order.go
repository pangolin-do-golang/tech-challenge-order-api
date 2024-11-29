package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge-order-api/internal/adapters/rest/controller"
	"github.com/pangolin-do-golang/tech-challenge-order-api/internal/core/order"
)

func RegisterOrderHandlers(router *gin.Engine, service order.IOrderService) {
	orderController := controller.NewOrderController(service)

	router.POST("/orders", orderController.Create)
	router.GET("/orders", orderController.GetAll)
	router.GET("/orders/:id", orderController.Get)
	router.PATCH("/orders/:id", orderController.Update)
}
