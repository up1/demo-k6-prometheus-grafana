package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type LoginRequest struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

func getUser(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	user := User{
		ID:   id,
		Name: "name",
	}
	return c.JSON(http.StatusOK, user)
}

func login(c echo.Context) error {
	req := new(LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	// Dummy authentication logic
	if req.User == "user" && req.Pass == "pass" {
		return c.JSON(http.StatusOK, map[string]string{"message": "login successful"})
	}
	return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
}

func main() {
	e := echo.New()
	e.GET("/users/:id", getUser)
	e.POST("/login", login)
	e.Logger.Fatal(e.Start(":8080"))
}
