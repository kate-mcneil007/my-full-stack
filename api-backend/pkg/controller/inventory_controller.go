package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/kate-mcneil007/my-full-stack/api-backend/pkg"
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

/*
You can declare methods with pointer receivers
This means the receiver type has the literal syntax *T for some type T (also, T cannot itself be a pointer such as *int.)
For example, the CreateInventoryItem method here is defined on *Controller
Pointer receivers can modify the value to which the receiver points (Controller)
So, our func points to struct Controller and passes in paramater var ctx of type echo.Context
The ctx parameter represents the context of the current HTTP request
error is the return type, the func will return an error
*/
func (c *Controller) CreateInventoryItem(ctx echo.Context) error {
	// Var body is a new instance of the Inventory struct from pkg
	body := pkg.Inventory{}

	// JSON data from the HTTP request body is decoded into a body variable & checks for decoding errors
	// If an error occurs, returns an HTTP response indicating a bad request along with the error message
	if err := json.NewDecoder(ctx.Request().Body).Decode(&body); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	/*
		This line is calling a method UpdateInventoryItem on the c.service object
		Passes two arguments:
			ctx.Request().Context()- This extracts context (information about the request) from incoming HTTP request, can be used for things like request cancellation, deadlines, and values associated with the request
			body
		Returns resp & err
	*/
	resp, err := c.service.UpdateInventoryItem(ctx.Request().Context(), body)
	// Erro handling
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	// No err then return 200 & response data
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
	body := pkg.Inventory{}

	/*
		The decoder is initialized to read from the Body of the HTTP request, which contains the JSON data sent by the client
		".Decode(&body)" attempts to decode the JSON data from the HTTP request's body
		into the body variable (where you want to store the decoded data)
		&body notation is used to pass the memory address of the body variable so that the
		decoder can populate it.
	*/
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
