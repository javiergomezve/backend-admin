package controllers

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/javiergomezve/backend-admin/database"
	"github.com/javiergomezve/backend-admin/models"
)

func AllProducts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 4
	offset := (page - 1) * limit
	var total int64

	var products []models.Product

	database.DB.Offset(offset).Limit(limit).Find(&products)

	database.DB.Model(&models.Product{}).Count(&total)

	return c.JSON(fiber.Map{
		"data": products,
		"meta": fiber.Map{
			"page":      page,
			"total":     total,
			"last_page": math.Ceil(float64((int(total) / limit))),
		},
	})
}

func CreateProduct(c *fiber.Ctx) error {
	var productDTO fiber.Map

	if err := c.BodyParser(&productDTO); err != nil {
		return err
	}

	product := models.Product{
		Title:       productDTO["title"].(string),
		Description: productDTO["description"].(string),
		Image:       productDTO["image"].(string),
		Price:       productDTO["price"].(float64),
	}

	database.DB.Create(&product)

	return c.JSON(product)
}

func GetProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	database.DB.Find(&product)

	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var productDTO fiber.Map

	if err := c.BodyParser(&productDTO); err != nil {
		return err
	}

	product := models.Product{
		Id:          uint(id),
		Title:       productDTO["title"].(string),
		Description: productDTO["description"].(string),
		Image:       productDTO["image"].(string),
		Price:       productDTO["price"].(float64),
	}

	database.DB.Model(&product).Updates(product)

	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	database.DB.Delete(&product)

	return c.JSON(nil)
}
