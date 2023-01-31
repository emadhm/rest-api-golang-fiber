package middleware

import (
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)


func CheckToken(c *fiber.Ctx) error {

	// Get the token from the request header
	tokenString := c.Get("Authorization")

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify that the token was signed with the correct secret key
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret key used for signing
		return []byte("secret_key"), nil
	})

	// Check if the token is valid
	if err != nil || !token.Valid {
		
		return c.Status(503).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	// Check if the token has expired
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expirationTime := int64(claims["exp"].(float64))
		if time.Now().Unix() > expirationTime {
			c.Status(503).JSON(fiber.Map{
			"message": "Token has Expired",
		})
		}
	}

	c.Next()

	return err
}
