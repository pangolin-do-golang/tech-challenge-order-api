package cart

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"os"
)

type Service struct {
	HttpClient resty.Client
}

func NewCartService() *Service {
	client := resty.New()
	client.SetBaseURL(os.Getenv("CART_SERVICE_URL"))
	return &Service{
		HttpClient: *client,
	}
}

func (s *Service) LoadCart(_ context.Context, clientID uuid.UUID) (*Cart, error) {
	var cart *Cart

	_, err := s.
		HttpClient.
		R().
		SetBody(LoadCardPayload{
			ClientID: clientID,
		}).
		SetResult(&cart).
		Post("/cart/loadcart")

	if err != nil {
		return nil, err
	}

	return cart, nil
}

type GetCartPayload struct {
	ClientID uuid.UUID `json:"client_id" binding:"required" format:"uuid"`
}

func (s *Service) GetFullCart(clientID uuid.UUID) (*Cart, error) {
	var cart *Cart
	_, err := s.
		HttpClient.
		R().
		SetBody(GetCartPayload{
			ClientID: clientID,
		}).
		SetResult(&cart).
		Post("/cart/overview")

	if err != nil {
		return nil, err
	}

	return cart, nil
}

type CleanupPayload struct {
	ClientID uuid.UUID `json:"client_id" binding:"required" format:"uuid"`
}

func (s *Service) Cleanup(clientID uuid.UUID) error {

	_, err := s.HttpClient.R().SetBody(CleanupPayload{
		ClientID: clientID,
	}).Post("/cart/cleanup")

	if err != nil {
		return err
	}

	return nil
}

type AddProductPayload struct {
	ClientID  uuid.UUID `json:"client_id" binding:"required" format:"uuid"`
	ProductID uuid.UUID `json:"product_id" binding:"required" format:"uuid"`
	Quantity  int       `json:"quantity" binding:"required,min=1" example:"1"`
	Comments  string    `json:"comments"`
}

func (s *Service) AddProduct(ctx context.Context, clientID uuid.UUID, product *Product) error {

	_, err := s.HttpClient.R().SetBody(AddProductPayload{
		ClientID:  clientID,
		ProductID: product.ProductID,
		Quantity:  product.Quantity,
		Comments:  product.Comments,
	}).Post("/cart/add-product")

	if err != nil {
		return err
	}

	return nil
}

type RemoveProductPayload struct {
	clientID  uuid.UUID `json:"client_id"`
	productID uuid.UUID `json:"product_id"`
}

type LoadCardPayload struct {
	ClientID uuid.UUID `json:"client_id" binding:"required" format:"uuid"`
}

func (s *Service) RemoveProduct(ctx context.Context, clientID uuid.UUID, productID uuid.UUID) error {
	_, err := s.HttpClient.R().SetBody(RemoveProductPayload{
		productID: productID,
		clientID:  clientID,
	}).Post("/cart/remove-product")

	if err != nil {
		return err
	}

	return nil
}

type EditProductPayload struct {
	ClientID  uuid.UUID `json:"client_id" binding:"required" format:"uuid"`
	ProductID uuid.UUID `json:"product_id" binding:"required" format:"uuid"`
	Quantity  int       `json:"quantity" binding:"required" example:"2"`
	Comments  string    `json:"comments"`
}

func (s *Service) EditProduct(ctx context.Context, clientID uuid.UUID, product *Product) error {

	_, err := s.
		HttpClient.
		R().
		SetBody(EditProductPayload{
			ClientID:  clientID,
			ProductID: product.ProductID,
			Quantity:  product.Quantity,
			Comments:  product.Comments,
		}).
		Post("/cart/edit-product")

	if err != nil {
		return nil
	}

	return nil
}

func (s *Service) GetProductByID(id uuid.UUID) (*Product, error) {
	var product Product

	_, err := s.HttpClient.R().
		SetPathParams(map[string]string{
			"id": id.String(),
		}).
		SetResult(&product).
		Get("/product/{id}")

	if err != nil {
		return nil, err
	}

	return &product, err
}
