package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/pangolin-do-golang/tech-challenge-order-api/docs"
	dbAdapter "github.com/pangolin-do-golang/tech-challenge-order-api/internal/adapters/db"
	"github.com/pangolin-do-golang/tech-challenge-order-api/internal/adapters/rest/server"
	"github.com/pangolin-do-golang/tech-challenge-order-api/internal/core/order"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title Tech Challenge Order Food API
// @version 0.1.0
// @description Fast Food API for FIAP Tech course

// @host localhost:8080
// @BasePath /
func main() {
	db, err := initDb()
	if err != nil {
		panic(err)
	}

	orderRepository := dbAdapter.NewPostgresOrderRepository(db)
	orderService := order.NewOrderService(orderRepository)

	restServer := server.NewRestServer(&server.RestServerOptions{
		OrderService: orderService,
	})

	restServer.Serve()
}

func initDb() (*gorm.DB, error) {
	_ = godotenv.Load()
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=America/Sao_Paulo",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	err = db.AutoMigrate(
		&dbAdapter.OrderPostgres{},
	)
	if err != nil {
		log.Fatalln(err)
	}

	return db, nil
}
