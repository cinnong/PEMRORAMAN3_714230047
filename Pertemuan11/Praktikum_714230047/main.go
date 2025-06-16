package main

import (
	"fmt"
	"inibackend/config"
	"inibackend/router"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func init () {
	// Load file .env saat program dijalankan
	err:= godotenv.Load()
	if err != nil{
		fmt.Println("Gagal memuat file .env:", err)
	}
}

func main() {
	app := fiber.New()

	//Logging Request
	app.Use(logger.New())

	//basic CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: strings.Join(config.GetAllOrigins(), ","),
		AllowCredentials: true,
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	//Route Mahasiswa/setup Router
	router.SetupRoutes(app)

	//Handler 404
	app.Use(func (c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": fiber.StatusNotFound,
			"Message": "Endpoint tidak ditemukan",
		})
		
	})

	//Baca PORT yang ada di .env
	port := os.Getenv("PORT")
	if port == ""{
		port = "3000" //Default port kalau tidak ada .env
	}

	//utk log cek konek di port mana
	log.Printf("Server running on port %s", port)
	if err := app.Listen(":" + port); err != nil{
		log.Fatal("Error starting server: %v",err)
	}//koneksi trptus

}


