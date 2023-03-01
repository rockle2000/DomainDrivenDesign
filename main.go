package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"DDD_Project/domain/repository"
	"DDD_Project/domain/service"
	"DDD_Project/infrastructure/config"
	"DDD_Project/infrastructure/persistence/datastore"
	"DDD_Project/infrastructure/transport/http"
)

func main() {
	cf, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalln("error while loading config ", err)
	}
	dbConfig := cf.GetDbConfig()
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DbName, dbConfig.SslMode))
	if err != nil {
		log.Fatalln("Error while connect to database", err)
	}
	customerDataStore := datastore.NewCustomerDatastore(db)
	customerRepo := repository.NewCustomerRepository(customerDataStore)
	customerService := service.NewCustomerService(customerRepo)
	customerTransport := http.NewCustomerHandler(customerService)

	e := echo.New()
	e.GET("/customer", customerTransport.GetListCustomer)
	e.GET("/customer/:id", customerTransport.GetCustomer)
	e.Logger.Fatal(e.Start(":8081"))

}
