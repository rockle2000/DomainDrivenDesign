package router

import (
	"DDD_Project/domain/repository"
	"DDD_Project/domain/service"
	"DDD_Project/infrastructure/middlewares"
	"DDD_Project/infrastructure/persistence/datastore"
	"DDD_Project/infrastructure/transport/http"
)

func (r *AppRouter) InitCustomerRouter() {
	db := r.db
	cf := r.config
	echoRouter := r.echo

	customerDataStore := datastore.NewCustomerDatastore(db)
	customerRepo := repository.NewCustomerRepository(customerDataStore)
	customerService := service.NewCustomerService(customerRepo)
	customerTransport := http.NewCustomerHandler(customerService)

	customerGroup := echoRouter.Group("customers")
	customerGroup.Use(middlewares.JWTAuth(cf))
	{
		customerGroup.GET("", customerTransport.GetListCustomer)
		customerGroup.GET("/:id", customerTransport.GetCustomer)
	}
}
