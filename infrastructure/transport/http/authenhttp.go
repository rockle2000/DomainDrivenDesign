package http

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"DDD_Project/domain/service"
)

type AuthenticationHandler struct {
	service service.AuthenticationService
}

func NewAuthenticationHandler(service service.AuthenticationService) *AuthenticationHandler {
	return &AuthenticationHandler{
		service: service,
	}
}

func (ah *AuthenticationHandler) Login(c echo.Context) error {
	var loginReq service.LoginReq
	if err := (&echo.DefaultBinder{}).BindBody(c, &loginReq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	res, err := ah.service.Login(loginReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": res,
	})
}

func (ah *AuthenticationHandler) Register(c echo.Context) error {
	var registerReq service.RegisterReq
	if err := (&echo.DefaultBinder{}).BindBody(c, &registerReq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err := ah.service.Register(registerReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Register successfully",
	})
}
