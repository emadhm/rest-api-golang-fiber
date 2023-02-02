package main

import (
	"emad.com/config"
	"emad.com/routes"
	"github.com/gofiber/fiber/v2"
    
    "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
    
    app := fiber.New()

    app.Use(cors.New())

    config.ConnectDB()
    
    routes.Routers(app)


   
    app.Listen(":3000")
}