package main

import "net/http"

func createRecipe(c echo.Context) error {
	return c.String(http.StatusOK, "Create a new recipe")
}

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
