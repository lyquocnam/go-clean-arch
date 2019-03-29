package main

import (
	"clean/handler"
	"clean/model"
	"clean/repository"
	"clean/storage"
	"clean/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

func main() {
	r := gin.Default()

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=clean password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(model.Customer{}, model.Email{})
	db.LogMode(true)

	// create customer storage
	customerStorage := storage.NewCustomerStorage(db)
	// create email storage
	emailStorage := storage.NewEmailStorage(db)

	// create customer repository
	customerRepo := repository.NewCustomerRepository(customerStorage)
	// create email repository
	emailRepo := repository.NewEmailRepository(emailStorage)

	// create use case
	customerUseCase := usecase.NewCustomerUseCase(customerRepo, emailRepo, time.Duration(10 * time.Second))

	// create handler
	handler.NewCustomerHandler(r, customerUseCase)

	// start app
	log.Fatal(r.Run(":7000"))
}


