package handler

import (
	"inibackend/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Homepage(c *fiber.Ctx) error {
	return c.SendString("Welcome to the homepage!")
}

func GetAllMahasiswa(c *fiber.Ctx) error{
	data, err := repository.GetAllMahasiswa(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengambil data mahasiswa", //gagal mengambil data dari database
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status" : fiber.StatusOK, //status 200
		"message" : "Data Mahasiswa berhasil diambil",
		"data":data,
	})
}


func GetMahasiswaByNPM(c *fiber.Ctx) error {
	npmStr := c.Params("npm")//Mengambil npm dari parameter url

	npm, err := strconv.Atoi(npmStr)//Mengubah string ke int
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "NPM harus berupa angka", //NPM tidak valid
		})
	}
	mhs, err := repository.GetMahasiswaByNPM(c.Context(), npm)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if mhs == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data mahasiswa tidak ditemukan", //Data mahasiswa tidak ditemukan
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK, //status 200
		"message": "Berhasil mengambil data mahasiswa", //Berhasil mengambil data dari database
		"data":    mhs, //data mahasiswa
	})
}