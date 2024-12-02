package cart

import (
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

type RemoveProductPayload struct {
	clientID  uuid.UUID `json:"client_id"`
	productID uuid.UUID `json:"product_id"`
}

func (s *Service) GetProductByID(id uuid.UUID) (*Product, error) {
	var product *Product

	_, err := s.HttpClient.R().
		SetPathParams(map[string]string{
			"id": id.String(),
		}).
		SetResult(&product).
		Get("/product/{id}")

	if err != nil {
		return nil, err
	}

	return product, nil
}
