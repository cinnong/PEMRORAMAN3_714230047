package router

import (
	"inibackend/config/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api:=app.Group("/api")

	api.Get("/", handler.Homepage)
	api.Get("/mahasiswa", handler.GetAllMahasiswa)
	api.Get("/mahasiswa/:npm", handler.GetMahasiswaByNPM)


}
