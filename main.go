package main

import (
	"go-test/api/controller"
	"go-test/api/router"
	"go-test/config"
	"go-test/database"
	"go-test/model"
	"go-test/repo"
	"go-test/usecase"
	"log"
	"net/http"
	"time"

	"gorm.io/gorm"
)

func main() {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	//Database
	db := database.ConnectionDB(&loadConfig)

	errMigrate := db.AutoMigrate(&model.Cars{}, &model.Orders{})
	if err != nil {
		log.Fatal("🚀 Could not DB Migrate", errMigrate)
	}

	//
	startServer(db)
}

func startServer(db *gorm.DB) {
	carRepo := repo.NewCarRepository(db)
	carUseCase := usecase.NewCarUseCase(carRepo)
	carController := controller.NewCarController(carUseCase)

	routes := router.NewServer(carController)

	server := &http.Server{
		Addr:           ":3400",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
