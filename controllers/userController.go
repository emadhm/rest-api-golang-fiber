package controllers

import (
	"emad.com/config"
	"emad.com/models"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Request struct {
	Name     string
	Email    string
	Role    string
	Password string
}

 func Create (c *fiber.Ctx) error {

	    var req = new(Request)

		if err := c.BodyParser(req); err != nil {   
			return err	
		}

		if req.Name == "" || req.Role == "" || req.Password == "" {
			return fiber.NewError(500, "all data required")
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

		if err != nil {
			return err
		}
		
		user := models.Users{
			Name: req.Name,
			Email: req.Email,
			Role: req.Role,
			Password: string(hash),
		}

		if err := config.DB.Create(&user).Error; err != nil {

		return err
	    }

		return c.JSON(fiber.Map{
			"message": "1 row add successfully",
			"user": user,
		})
	}

	func List (c *fiber.Ctx) error {
	    
		var user = new([]models.Users)

		if err := config.DB.Find(&user).Error; err != nil {

		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	    }


		return c.JSON(fiber.Map{
			"message": "list of users",
			"users": user,
		})
		
	}

	func One (c *fiber.Ctx) error {
	    
		var user = new(models.Users)
	
		 id :=c.Params("id")

		if err := config.DB.First(&user, id).Error; err !=nil {
			if err == gorm.ErrRecordNotFound {
				return c.Status(501).JSON(fiber.Map{
			"message": "user not found",
		 })
			}

			return err
		} 

		
		return c.JSON(fiber.Map{
			"message": "one user",
			"user": user,
		})
		
	}


	func Update (c *fiber.Ctx) error {
	    
		var req = new(Request)
	     user := new(models.Users)
	
		 id :=c.Params("id")

		  err := config.DB.First(&user, id).Error
			
		 if err == gorm.ErrRecordNotFound {
				return err 
			}

			if err := c.BodyParser(&req); err != nil {
		     return err
		}
		if req.Name == "" || req.Role == "" || req.Password == "" {
			return fiber.NewError(500, "all data required")
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

		if err != nil {
			return err
		}
        
		user.Name = req.Name
		user.Role = req.Role
		user.Password = string(hash)
		

	if  config.DB.Save(&user).RowsAffected == 0 {
		return c.Status(503).JSON(fiber.Map{
			"message": "update falled",
		})
	}

			
		
		return c.JSON(fiber.Map{
			"message": "One User updated",
			"user": user,
		})
		
	}

	
	


	func Delete (c *fiber.Ctx) error {
	    
		var user = new(models.Users)
	
		 id :=c.Params("id")

		 err := config.DB.First(&user, id).Error
			
		 if err == gorm.ErrRecordNotFound {

				return c.Status(501).JSON(fiber.Map{
			"message": "user does not exsist",
		 })
			}


	 if  config.DB.Delete(&user).RowsAffected == 0 {
		return c.Status(503).JSON(fiber.Map{
			"message": "delete falled",
		})
	}

			
		
		return c.JSON(fiber.Map{
			"message": "one user deleted",
			"user": user,
		})
		
	}