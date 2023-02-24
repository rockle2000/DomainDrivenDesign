package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"DDD_Project/domain/repository"
	"DDD_Project/domain/service"
	"DDD_Project/infrastructure/persistence/datastore"
	"DDD_Project/infrastructure/transport/http"
)

func main() {
	port := "5432"
	host := "localhost"
	user := "postgres"
	password := "123456"
	dbname := "test"
	sslmode := "disable"

	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode))
	if err != nil {
		log.Fatalln(err)
	}
	customerDataStore := datastore.NewCustomerDatastore(db)
	customerRepo := repository.NewCustomerRepository(customerDataStore)
	customerService := service.NewCustomerService(customerRepo)
	customerTransport := http.NewCustomerHandler(customerService)

	var r = gin.Default()
	r.GET("/customer", customerTransport.GetListCustomer)
	log.Fatalln(r.Run(":8081"))

}
