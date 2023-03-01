package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"DDD_Project/domain/service"
)

type CustomerHandler struct {
	customerService service.CustomerService
}

func NewCustomerHandler(service service.CustomerService) *CustomerHandler {
	return &CustomerHandler{
		customerService: service,
	}
}

func (ch *CustomerHandler) GetCustomer(c echo.Context) error {
	id := c.Param("id")
	list, err := ch.customerService.GetCustomer(id)
	if err != nil {
		fmt.Println("er", err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})

	}
	return c.JSON(http.StatusOK, list)
}

func (ch *CustomerHandler) GetListCustomer(c echo.Context) error {
	list, err := ch.customerService.GetAllCustomer()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": list,
	})
}
