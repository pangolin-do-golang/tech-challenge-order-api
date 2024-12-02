package cart

import "github.com/google/uuid"

type Cart struct {
	ID       uuid.UUID  `json:"id"`
	ClientID uuid.UUID  `json:"client_id"`
	Products []*Product `json:"products"`
}

type Product struct {
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Comments  string    `json:"comments,omitempty"`
	Price     float64   `json:"price"`
}
