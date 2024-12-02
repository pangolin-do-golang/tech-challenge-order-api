package server_test

import (
	"github.com/pangolin-do-golang/tech-challenge-order-api/internal/adapters/rest/server"
	"github.com/pangolin-do-golang/tech-challenge-order-api/mocks"
	"testing"
)

func TestServeStartsServerSuccessfully(t *testing.T) {
	service := new(mocks.IOrderService)
	rs := server.NewRestServer(&server.RestServerOptions{
		OrderService: service,
	})

	go func() {
		rs.Serve()
	}()
}
