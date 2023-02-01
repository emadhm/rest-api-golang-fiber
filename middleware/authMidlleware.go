package middleware

import (
	"fmt"
	"os"
	"strings"
    "github.com/joho/godotenv"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)


func CheckToken(c *fiber.Ctx) error {
	
	
	err := godotenv.Load()
    if err != nil {
	return err
    }
	secretKey := os.Getenv("SECRET_KEY")

	// Get the token from the request header
	tokenString := c.Get("Authorization")
	if len(tokenString) < 7 || strings.ToLower(tokenString[0:7]) != "bearer " {
		c.Status(501).JSON(fiber.Map{
			"messege": "Authorization header format must be 'Bearer [token]'",
		})
		
	}

	// Remove the "bearer" keyword from the token string
	 tokenString = tokenString[7:]

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify that the token was signed with the correct secret key and check expire token
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret key used for signing
		return []byte(secretKey), nil
	})

	// Check if the token is valid
	if err != nil || !token.Valid {
		
		return err
	}

	c.Next()

	return err
}
