package handler

//package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRoutes() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/hello", hello)
	e.POST("/inventory", createInventoryItem)
	e.GET("/inventory/:id", getInventoryItem)
	e.PUT("/inventory", updateInventoryItem)
	e.DELETE("/inventory", deleteInventoryItem)
	e.Logger.Fatal(e.Start(":3000"))
}
