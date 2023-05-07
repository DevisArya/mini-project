package middleware

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var secretKey = func() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	key := os.Getenv("SECRET_KEY")

	return key
}()

var IsloggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(secretKey),
})

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		isAdmin := claims["role"]

		if isAdmin == false {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

func UserActValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return echo.ErrBadRequest
		}

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		claimsID := fmt.Sprint(claims["userId"])
		convClaimsID, ok := strconv.Atoi(claimsID)

		if ok != nil {
			panic("invalid type")
		}

		if int(convClaimsID) != id {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}
