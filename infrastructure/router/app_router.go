package router

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	"DDD_Project/infrastructure/config"
)

type AppRouter struct {
	echo   *echo.Echo
	db     *sqlx.DB
	config *config.AppConfig
}

func NewAppRouter(e *echo.Echo, db *sqlx.DB, config *config.AppConfig) *AppRouter {
	return &AppRouter{
		echo:   e,
		db:     db,
		config: config,
	}
}

func InitRouter(router *AppRouter) {
	router.InitAuthenticationRouter()
	router.InitCustomerRouter()
}
