package http

import (
	"net/http"
	"strconv"

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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	list, err := ch.customerService.GetCustomer(id)
	if err != nil {
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
