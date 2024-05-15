package main

import (
	"github.com/manimovassagh/mani-app/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// User Handlers

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to the Recipe Sharing App!")
	})

	// User routes
	e.GET("/users", handlers.GetUsers)
	e.POST("/users", handlers.CreateUser)

	// Recipe routes
	e.GET("/recipes", handlers.GetRecipes)
	e.POST("/recipes", handlers.CreateRecipe)

	e.Logger.Fatal(e.Start(":8080"))
}
