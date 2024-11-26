package controller

import (
	"github.com/pangolin-do-golang/tech-challenge-order-api/internal/errutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge-order-api/internal/core/order"
)

type OrderController struct {
	AbstractController
	service order.IOrderService
}

func NewOrderController(service order.IOrderService) *OrderController {
	return &OrderController{
		service: service,
	}
}

// GetAll Get all order's list
// @Summary Get order list
// @Description Get all order's list
// @Tags Order
// @Accept json
// @Produce json
// @Success 200 {object} []order.Order{}
// @Success 500 {object} HTTPError
// @Router /orders [get]
func (ctrl *OrderController) GetAll(c *gin.Context) {
	orderSlice, err := ctrl.service.GetAll()

	if err != nil {
		ctrl.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, orderSlice)
}

type CreateOrderPayload struct {
	ClientID uuid.UUID `json:"client_id" binding:"required" format:"uuid"`
}

// Create Order godoc
// @Summary Create order from Cart
// @Description Create order from Cart
// @Param payload body controller.CreateOrderPayload true "CreateOrderPayload"
// @Tags Order
// @Accept  json
// @Produce  json
// @Success 200 {object} order.Order{}
// @Failure 400 {object} HTTPError
// @Failure 422 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /orders [post]
func (ctrl *OrderController) Create(c *gin.Context) {
	payload := &CreateOrderPayload{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		ctrl.Error(c, errutil.NewInputError(err))

		return
	}

	o, err := ctrl.service.Create(payload.ClientID)
	if err != nil {
		ctrl.Error(c, err)

		return
	}

	c.JSON(http.StatusOK, o)
}

// Get an order by ID
// @Summary Get order by ID
// @Description Get an order by ID
// @Tags Order
// @Param id path string true "ID of the order"
// @Accept json
// @Produce json
// @Success 200 {object} order.Order{}
// @Failure 400 {object} HTTPError
// @Failure 422 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /orders/{id} [get]
func (ctrl *OrderController) Get(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		ctrl.Error(c, errutil.NewInputError(err))
		return
	}

	o, err := ctrl.service.Get(id)
	if err != nil {
		ctrl.Error(c, err)

		return
	}

	c.JSON(http.StatusOK, o)
}

type UpdateOrderPayload struct {
	Status string `json:"status" binding:"required" example:"paid"`
}

// Update Order godoc
// @Summary Update an Order
// @Description Update by json an Order
// @Param id path string true "ID of the Order"
// @Param payload body controller.UpdateOrderPayload true "UpdateOrderPayload"
// @Tags Order
// @Accept  json
// @Produce  json
// @Success 200 {object} order.Order{}
// @Failure 400 {object} HTTPError
// @Failure 422 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /orders/{id} [patch]
func (ctrl *OrderController) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	payload := &UpdateOrderPayload{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		ctrl.Error(c, errutil.NewInputError(err))
		return
	}

	o, err := ctrl.service.Update(&order.Order{
		ID:     id,
		Status: strings.ToUpper(payload.Status),
	})
	if err != nil {
		ctrl.Error(c, err)
		return
	}

	c.JSON(http.StatusOK, o)
}
