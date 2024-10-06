package handlers

import (
	"RPN/cmd/models"
	"RPN/cmd/services"
	"RPN/cmd/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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

func Login(c echo.Context) error {
	var user User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}
	//var usermodel models.User
	usermodel,err := models.FindUserByUsername(user.Username)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid username or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(usermodel.Password), []byte(user.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid username or password")
	}
	tokenString,err := services.GenerateJWT(user.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to generate token")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"token": tokenString,
	})

}