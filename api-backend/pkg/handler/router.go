package handler

//package main

import (
	controller "github.com/kate-mcneil007/my-full-stack/api-backend/pkg/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRoutes(c controller.HandlerInterface) {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/hello", c.Hello)
	e.POST("/inventory", c.CreateInventoryItem)
	e.GET("/inventory/:id", c.GetInventoryItem)
	e.PUT("/inventory", c.UpdateInventoryItem)
	e.DELETE("/inventory", c.DeleteInventoryItem)
	e.Logger.Fatal(e.Start(":3000"))
}
