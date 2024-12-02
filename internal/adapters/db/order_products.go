package db

import (
	"context"
	"github.com/pangolin-do-golang/tech-challenge-order-api/internal/core/order"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostgresOrderProductsRepository struct {
	db *gorm.DB
}

type OrderProductPostgres struct {
	BaseModel
	OrderID   uuid.UUID     `gorm:"type:uuid"`
	ProductID uuid.UUID     `gorm:"type:uuid"`
	Quantity  int           `gorm:"quantity"`
	Comments  string        `gorm:"comments"`
	Order     OrderPostgres `gorm:"foreignKey:OrderID"`
}

func (op *OrderProductPostgres) TableName() string {
	return "order_products"
}

func NewPostgresOrderProductsRepository(db *gorm.DB) order.IOrderProductRepository {
	return &PostgresOrderProductsRepository{db: db}
}

func (p *PostgresOrderProductsRepository) Create(_ context.Context, orderID uuid.UUID, product *order.Product) error {
	orderProduct := OrderProductPostgres{
		OrderID:   orderID,
		ProductID: product.ProductID,
		Quantity:  product.Quantity,
		Comments:  product.Comments,
	}

	result := p.db.Create(&orderProduct)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
