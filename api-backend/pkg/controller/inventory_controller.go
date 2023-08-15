package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	database "github.com/kate-mcneil007/my-full-stack/api-backend/pkg/db"
	"github.com/kate-mcneil007/my-full-stack/api-backend/pkg/service"
	"github.com/labstack/echo/v4"
)

// HandlerInterface defines the methods that a handler should implement
type HandlerInterface interface {
	Hello(c echo.Context) error
	CreateInventoryItem(c echo.Context) error
	GetInventoryItem(c echo.Context) error
	UpdateInventoryItem(c echo.Context) error
	DeleteInventoryItem(c echo.Context) error
}

// Controller implements the HandlerInterface
type Controller struct {
	service service.ServiceInterface
	// logger ? maybe
}

// NewController creates a new instance of the Controller
func NewController(s *service.Service) HandlerInterface {
	return &Controller{
		service: s,
	}
}

// Implement methods for the HandlerInterface
func (c *Controller) Hello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}

func (c *Controller) CreateInventoryItem(ctx echo.Context) error {
	body := database.Inventory{}

	if err := json.NewDecoder(ctx.Request().Body).Decode(&body); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	resp, err := c.service.UpdateInventoryItem(ctx.Request().Context(), body)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(200, resp)
}

func (c *Controller) GetInventoryItem(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error")
	}
	resp, err := c.service.GetInventoryItem(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(500, err)
	}
	// Your logic to retrieve an inventory item using the provided ID
	return ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) UpdateInventoryItem(ctx echo.Context) error {
	body := database.Inventory{}

	if err := json.NewDecoder(ctx.Request().Body).Decode(&body); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	resp, err := c.service.UpdateInventoryItem(ctx.Request().Context(), body)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(200, resp)
}

// return bool
func (c *Controller) DeleteInventoryItem(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Error")
	}
	resp, err := c.service.GetInventoryItem(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(500, err)
	}
	// Your logic to retrieve an inventory item using the provided ID
	return ctx.JSON(http.StatusOK, resp)

}
