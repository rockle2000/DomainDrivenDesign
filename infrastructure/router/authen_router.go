package router

import (
	"DDD_Project/domain/repository"
	"DDD_Project/domain/service"
	"DDD_Project/infrastructure/persistence/datastore"
	"DDD_Project/infrastructure/transport/http"
)

func (r *AppRouter) InitAuthenticationRouter() {
	db := r.db
	cf := r.config
	echoRouter := r.echo

	customerDataStore := datastore.NewCustomerDatastore(db)
	authenticationRepo := repository.NewAuthenticationRepo(customerDataStore)
	authenticationService := service.NewAuthenticationService(authenticationRepo, cf)
	authenticationTransport := http.NewAuthenticationHandler(authenticationService)

	echoRouter.POST("/register", authenticationTransport.Register)
	echoRouter.POST("/login", authenticationTransport.Login)
}
