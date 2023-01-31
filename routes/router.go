package routes

import (
	"emad.com/auth"
	"emad.com/controllers"
	"emad.com/middleware"
	"github.com/gofiber/fiber/v2"
)


func Routers (app *fiber.App)  {

	 api := app.Group("/api")  // /api
	 login := api.Group("/login")
     user := api.Group("/user")  // /api/users


	 login.Post("",auth.Login)
	  
	 user.Get("/list",controllers.List)
	 user.Get("/one/:id", middleware.CheckToken, controllers.One)
 	 user.Post("/create",controllers.Create)
	 user.Put("/update/:id",controllers.Update)
	 user.Delete("/delete/:id",controllers.Delete)




	
}