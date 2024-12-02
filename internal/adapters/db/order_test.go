package db_test

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pangolin-do-golang/tech-challenge-order-api/mocks"
	"gorm.io/driver/postgres"
	"testing"

	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge-order-api/internal/adapters/db"
	"github.com/pangolin-do-golang/tech-challenge-order-api/internal/core/order"
	"github.com/pangolin-do-golang/tech-challenge-order-api/internal/errutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestCreateOrderSuccessfully(t *testing.T) {
	mockDB := new(mocks.IDB)
	mockDB.On("Create", mock.Anything).Return(&gorm.DB{})

	repo := db.NewPostgresOrderRepository(mockDB)
	newOrder, err := repo.Create(&order.Order{ClientID: uuid.New(), Status: "created", TotalAmount: 100.0})

	assert.NoError(t, err)
	assert.NotNil(t, newOrder)
}

func TestCreateOrderFails(t *testing.T) {
	mockDB := new(mocks.IDB)
	mockDB.On("Create", mock.Anything).Return(&gorm.DB{Error: errors.New("create error")})

	repo := db.NewPostgresOrderRepository(mockDB)
	newOrder, err := repo.Create(&order.Order{ClientID: uuid.New(), Status: "created", TotalAmount: 100.0})

	assert.Error(t, err)
	assert.Nil(t, newOrder)
}

func TestGetOrderSuccessfully(t *testing.T) {
	mockDB := new(mocks.IDB)
	mockDB.On("First", mock.Anything, mock.Anything, mock.Anything).Return(&gorm.DB{})

	repo := db.NewPostgresOrderRepository(mockDB)
	o, err := repo.Get(uuid.New())

	assert.NoError(t, err)
	assert.NotNil(t, o)
}

func TestGetOrderNotFound(t *testing.T) {
	mockDB := new(mocks.IDB)
	mockDB.On("First", mock.Anything, mock.Anything, mock.Anything).Return(&gorm.DB{Error: gorm.ErrRecordNotFound})

	repo := db.NewPostgresOrderRepository(mockDB)
	o, err := repo.Get(uuid.New())

	assert.ErrorIs(t, err, errutil.ErrRecordNotFound)
	assert.Nil(t, o)
}

func TestGetAllOrdersSuccessfully(t *testing.T) {
	d, m, err := sqlmock.New()
	conn, err := gorm.Open(postgres.New(postgres.Config{Conn: d, DriverName: "postgres"}))

	m.ExpectQuery("SELECT .+").WillReturnRows(sqlmock.NewRows([]string{"id", "client_id", "total_amount", "status", "customer_id", "product_id"}))
	repo := db.NewPostgresOrderRepository(conn)
	_, err = repo.GetAll()

	assert.NoError(t, err)
}
