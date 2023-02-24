package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

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

func (ch *CustomerHandler) GetCustomer(c *gin.Context) {
	id := c.Param("id")
	list, err := ch.customerService.GetCustomer(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": list,
	})
}

func (ch *CustomerHandler) GetListCustomer(c *gin.Context) {
	list, err := ch.customerService.GetAllCustomer()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": list,
	})
}
