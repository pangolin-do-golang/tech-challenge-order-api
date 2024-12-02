package cart_test

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge-order-api/internal/core/cart"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func createMockedProduct() cart.Product {
	return cart.Product{
		ProductID: uuid.New(),
		Quantity:  1,
		Comments:  "Mocked Product",
		Price:     10.0,
	}
}

func setupMockServer(responseBody interface{}, statusCode int) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)

		err := json.NewEncoder(w).Encode(responseBody)

		if err != nil {
			panic(err)
		}
	})
	return httptest.NewServer(handler)
}

func TestNewCartService(t *testing.T) {
	os.Setenv("CART_SERVICE_URL", "http://mocked.com")
	service := cart.NewCartService()
	assert.Equal(t, "http://mocked.com", service.HttpClient.BaseURL)
}

func TestService_GetFullCart(t *testing.T) {
	mockCart := cart.Cart{ /* fill with mock data if available */ }
	mockServer := setupMockServer(mockCart, http.StatusOK)
	defer mockServer.Close()

	os.Setenv("CART_SERVICE_URL", mockServer.URL)
	service := cart.NewCartService()

	clientID := uuid.New()
	cart, err := service.GetFullCart(clientID)

	require.NoError(t, err)
	assert.Equal(t, &mockCart, cart)
}

func TestService_Cleanup(t *testing.T) {
	mockServer := setupMockServer(nil, http.StatusOK)
	defer mockServer.Close()

	os.Setenv("CART_SERVICE_URL", mockServer.URL)
	service := cart.NewCartService()

	clientID := uuid.New()
	err := service.Cleanup(clientID)

	require.NoError(t, err)
}

func TestService_GetProductByID(t *testing.T) {
	mockProduct := createMockedProduct()
	mockServer := setupMockServer(mockProduct, http.StatusOK)
	defer mockServer.Close()

	os.Setenv("CART_SERVICE_URL", mockServer.URL)
	service := cart.NewCartService()

	productID := uuid.New()
	product, err := service.GetProductByID(productID)

	require.NoError(t, err)
	assert.Equal(t, &mockProduct, product)
}
