package handlers

import (
	"RPN/cmd/models"
	"RPN/cmd/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)
	

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func Registration(c echo.Context) error {
	var user User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	if !utils.ValidatePassword(user.Password) {
		return c.JSON(http.StatusBadRequest, "password invalid")
	}

	userexist,err := models.UsernameExist(user.Username)
	if err  != nil || userexist{
		return c.JSON(http.StatusBadRequest, "Username already used")
	}else if !userexist {
		err = models.CreateUser(user.Username, user.Password)
		if err != nil {
            return c.JSON(http.StatusInternalServerError, "Failed to create user")
        }
	} 
	return c.JSON(http.StatusCreated, "User registered successfully")
}