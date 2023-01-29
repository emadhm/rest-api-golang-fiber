package routes

import (
     "github.com/gofiber/fiber/v2"
	 "emad.com/controllers"

)


func Routers (app *fiber.App)  {

	 api := app.Group("/api")  // /api
     user := api.Group("/user")  // /api/users

	  
	 user.Get("/list",controllers.List)
	 user.Get("/one/:id",controllers.One)
 	 user.Post("/create",controllers.Create)
	 user.Put("/update/:id",controllers.Update)
	 user.Delete("/delete/:id",controllers.Delete)




	
}