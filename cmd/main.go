package main

import (
	"RPN/cmd/db"
	"RPN/cmd/handlers"
	"RPN/cmd/services"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)


func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*services.JwtCustomClaims)
	name := claims.Username
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func main() {
	e := echo.New()
	
	db.Init()

	e.Use(middleware.Logger())
	//e.Use(middleware.Recover())
	// e.Use(echojwt.WithConfig(echojwt.Config{
	// 	// ...
	// 	SigningKey:             []byte(services.JwtSecret),
	// 	// ...
	// }))

	e.POST("/login", handlers.Login)
	e.POST("/singup", handlers.Registration)
	e.GET("/", accessible)

	// Restricted group
	r := e.Group("/api")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(services.JwtCustomClaims)
		},
		SigningKey: []byte(services.JwtSecret),
	}
	r.Use(echojwt.WithConfig(config))
	r.GET("", restricted)
	e.Logger.Fatal(e.Start(":8000"))
}