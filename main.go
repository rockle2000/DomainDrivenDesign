package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"

	"DDD_Project/infrastructure/config"
	"DDD_Project/infrastructure/router"
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

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//e.Use(middlewares.MiddlewareLogging)
	r := router.NewAppRouter(e, db, cf)

	router.InitRouter(r)
	e.Logger.Fatal(e.Start(":8081"))

}
