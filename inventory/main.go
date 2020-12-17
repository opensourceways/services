package main

import (
	"inventory/handler"
	"inventory/model"
	pb "inventory/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("inventory"),
		service.Version("latest"),
	)

	// Connect to the database
	db, err := gorm.Open(postgres.Open("postgresql://postgres@localhost:5432/inventory?sslmode=disable"), nil)
	if err != nil {
		logger.Fatal(err)
	}
	if err := db.AutoMigrate(&model.Order{}, &model.Item{}, &model.Topup{}); err != nil {
		logger.Fatal(err)
	}

	// Register handler
	pb.RegisterInventoryHandler(srv.Server(), &handler.Inventory{DB: db})

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
