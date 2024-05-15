package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// User Handlers
func getUsers(c echo.Context) error {
	return c.String(http.StatusOK, "Get all users")
}

func createUser(c echo.Context) error {
	return c.String(http.StatusOK, "Create a new user")
}

// Recipe Handlers
func getRecipes(c echo.Context) error {
	return c.String(http.StatusOK, "Get all recipes")
}

func createRecipe(c echo.Context) error {
	return c.String(http.StatusOK, "Create a new recipe")
}

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
	e.GET("/users", getUsers)
	e.POST("/users", createUser)

	// Recipe routes
	e.GET("/recipes", getRecipes)
	e.POST("/recipes", createRecipe)

	e.Logger.Fatal(e.Start(":8080"))
}
