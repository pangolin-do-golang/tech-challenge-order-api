package controller_test

import (
	"github.com/pangolin-do-golang/tech-challenge-order-api/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge-order-api/internal/adapters/rest/controller"
	"github.com/pangolin-do-golang/tech-challenge-order-api/internal/core/order"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateOrderWithValidPayload(t *testing.T) {
	mockService := new(mocks.IOrderService)
	mockService.On("Create", mock.Anything).Return(&order.Order{}, nil)
	router := gin.Default()
	ctrl := controller.NewOrderController(mockService)
	router.POST("/orders", ctrl.Create)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/orders", strings.NewReader(`{"client_id":"550e8400-e29b-41d4-a716-446655440000"}`))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateOrderWithInvalidPayload(t *testing.T) {
	mockService := new(mocks.IOrderService)
	router := gin.Default()
	ctrl := controller.NewOrderController(mockService)
	router.POST("/orders", ctrl.Create)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/orders", strings.NewReader(`{"client_id":"invalid-uuid"}`))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetAllOrdersSuccessfully(t *testing.T) {
	mockService := new(mocks.IOrderService)
	mockService.On("GetAll").Return([]order.Order{}, nil)
	router := gin.Default()
	ctrl := controller.NewOrderController(mockService)
	router.GET("/orders", ctrl.GetAll)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/orders", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetOrderWithValidID(t *testing.T) {
	mockService := new(mocks.IOrderService)
	mockService.On("Get", mock.Anything).Return(&order.Order{}, nil)
	router := gin.Default()
	ctrl := controller.NewOrderController(mockService)
	router.GET("/orders/:id", ctrl.Get)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/orders/550e8400-e29b-41d4-a716-446655440000", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetOrderWithInvalidID(t *testing.T) {
	mockService := new(mocks.IOrderService)
	router := gin.Default()
	ctrl := controller.NewOrderController(mockService)
	router.GET("/orders/:id", ctrl.Get)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/orders/invalid-uuid", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateOrderWithValidPayload(t *testing.T) {
	mockService := new(mocks.IOrderService)
	mockService.On("Update", mock.Anything).Return(&order.Order{}, nil)
	router := gin.Default()
	ctrl := controller.NewOrderController(mockService)
	router.PATCH("/orders/:id", ctrl.Update)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/orders/550e8400-e29b-41d4-a716-446655440000", strings.NewReader(`{"status":"paid"}`))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateOrderWithInvalidPayload(t *testing.T) {
	mockService := new(mocks.IOrderService)
	router := gin.Default()
	ctrl := controller.NewOrderController(mockService)
	router.PATCH("/orders/:id", ctrl.Update)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/orders/550e8400-e29b-41d4-a716-446655440000", strings.NewReader(`{"status":""}`))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
