package controllers

import (
	"net/http"
	"prakerja6/config"
	"prakerja6/models"

	"github.com/labstack/echo/v4"
)

func CreatePacketController(c echo.Context) error {
	packet := models.Packet{}
	c.Bind(&packet)

	result := config.DB.Create(&packet)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Gagal memasukkan data", Status: false, Data: nil,
		})
	}
	return c.JSON(http.StatusCreated, models.BaseResponse{
		Message: "Success memasukkan data", Status: true, Data: packet,
	})
}

func PacketController(c echo.Context) error {
	var packets []models.Packet

	result := config.DB.Find(&packets)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: "Gagal mendapatkan data", Status: false, Data: nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Success mendapatkan data", Status: true, Data: packets,
	})
}
