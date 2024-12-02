package order

import (
	"context"
	"errors"

	"time"

	"github.com/google/uuid"
)

const (
	StatusPending   = "PENDING"
	StatusCreated   = "CREATED"
	StatusPreparing = "PREPARING"
	StatusReady     = "READY"
	StatusFinished  = "FINISHED"
	StatusPaid      = "PAID"
	StatusDeclined  = "DECLINED"
	StatusCanceled  = "CANCELED"
)

var ErrInvalidStatus = errors.New("invalid status transition")

type Order struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	ClientID    uuid.UUID `json:"client_id"`
	TotalAmount float64   `json:"total_amount"`
	Status      string    `json:"status"`
}

func (o Order) ValidateStatusTransition(nextStatus string) error {
	switch o.Status {
	case StatusCreated:
		if nextStatus != StatusPending {
			return ErrInvalidStatus
		}
	case StatusPending:
		if nextStatus != StatusPaid && nextStatus != StatusDeclined {
			return ErrInvalidStatus
		}
	case StatusPaid:
		if nextStatus != StatusPreparing {
			return ErrInvalidStatus
		}
	case StatusPreparing:
		if nextStatus != StatusFinished {
			return ErrInvalidStatus
		}
	case StatusCanceled:
		return ErrInvalidStatus
	default:
		return ErrInvalidStatus
	}
	return nil
}

type Product struct {
	ClientID  uuid.UUID `json:"client_id"`
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Comments  string    `json:"comments"`
	Total     float64   `json:"total"`
}

type IOrderService interface {
	Get(id uuid.UUID) (*Order, error)
	Create(clientID uuid.UUID) (*Order, error)
	GetAll() ([]Order, error)
	Update(order *Order) (*Order, error)
}

type IOrderProductRepository interface {
	Create(ctx context.Context, orderID uuid.UUID, product *Product) error
}

type IOrderRepository interface {
	Create(order *Order) (*Order, error)
	Update(order *Order) error
	Get(id uuid.UUID) (*Order, error)
	GetAll() ([]Order, error)
}
