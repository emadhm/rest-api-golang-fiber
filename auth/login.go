package auth

import (
	"os"
	"time"

	"emad.com/config"
	"emad.com/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {

Email string
Password string
}

func Login (c *fiber.Ctx) error {

	var user = new(models.Users)

	var req = new(LoginRequest)

		if err := c.BodyParser(req); err != nil {   
			return err	
		}

		if req.Email == "" || req.Password == "" {
			return fiber.NewError(fiber.StatusBadRequest, "invalid login data")
		}

		 if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
				return c.Status(501).JSON(fiber.Map{
			"message": "user not found",
		 })
			}

			return err
		}

       if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return  c.Status(501).JSON(fiber.Map{
			"message": "incorrect email or password",
		 })
      }

      token, err := generateJWT(int(user.ID))
		if err != nil {
			return err
		}


	  checkAuth(int(user.ID),token)

   return c.JSON(fiber.Map{
			"message": "loged in seccessfully",
			"user": user,
			"token":token,
		 })
		

}

func generateJWT(id int) (string, error) {

	err := godotenv.Load()
	if err != nil {
		return "",err
  	 }
     secretKey := os.Getenv("SECRET_KEY")


	// Create JWT token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		"exp":      time.Now().Add(time.Hour * 720).Unix(),
	})

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func checkAuth (id int, token string) error {

	 var auth = new(models.Auths)
	
   if err := config.DB.Where("user_id = ?", id).First(&auth).Error; err != nil {

		if err == gorm.ErrRecordNotFound {

			auth.User_id = id
			auth.Token = token

			if err := config.DB.Create(&auth).Error; err != nil {

		     return err
	        }
		}
	    	
		return err
	}

	 auth.User_id = id
	 auth.Token = token

    if err := config.DB.Save(&auth).Error; err!=nil {
		 return err
	}
	 
	return nil

}