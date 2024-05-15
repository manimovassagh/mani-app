package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateRecipe(c echo.Context) error {
	return c.String(http.StatusOK, "Create a new recipe")
}

func GetUsers(c echo.Context) error {
	return c.String(http.StatusOK, "Get all users")
}

func CreateUser(c echo.Context) error {
	return c.String(http.StatusOK, "Create a new user")
}

// GetRecipes Recipe Handlers
func GetRecipes(c echo.Context) error {
	return c.String(http.StatusOK, "Get all recipes")
}
