package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge-order-api/internal/adapters/rest/handler"
	"github.com/pangolin-do-golang/tech-challenge-order-api/internal/adapters/rest/middleware"
	"github.com/pangolin-do-golang/tech-challenge-order-api/internal/core/order"
)

type RestServer struct {
	orderService order.IOrderService
}

type RestServerOptions struct {
	OrderService order.IOrderService
}

func NewRestServer(options *RestServerOptions) *RestServer {
	return &RestServer{
		orderService: options.OrderService,
	}
}

func (rs RestServer) Serve() {
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())
	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	handler.RegisterOrderHandlers(r, rs.orderService)
	handler.RegisterSwaggerHandlers(r)
	err := r.Run("0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
}
