package handler

import (
	"clean/usecase"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type customerHandler struct {
	router *gin.Engine
	customerUseCase usecase.CustomerUseCase
}

func NewCustomerHandler(router *gin.Engine, customerUseCase usecase.CustomerUseCase) *customerHandler {
	handler := &customerHandler{
		router: router,
		customerUseCase: customerUseCase,
	}

	// map routes
	handler.router.GET("/:id", handler.GetCustomerByIdHandler)

	return handler
}

type CustomerHandler interface {
	GetCustomerByIdHandler(c *gin.Context)
}

func (s *customerHandler) GetCustomerByIdHandler(c *gin.Context) {
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id is required",
		})
		return
	}

	customer, err := s.customerUseCase.GetCustomerById(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, customer)
}