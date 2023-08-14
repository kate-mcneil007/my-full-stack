package handler

//package main

import (
	"github.com/kate-mcneil007/my-full-stack/api-backend/pkg/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRoutes(c controller.Controller) {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/hello", c.hello)
	e.POST("/inventory", c.createInventoryItem)
	e.GET("/inventory/:id", c.getInventoryItem)
	e.PUT("/inventory", c.updateInventoryItem)
	e.DELETE("/inventory", c.deleteInventoryItem)
	e.Logger.Fatal(e.Start(":3000"))
}
