package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/javiergomezve/backend-admin/database"
	"github.com/javiergomezve/backend-admin/middlewares"
	"github.com/javiergomezve/backend-admin/models"
)

func AllProducts(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorize(c, "products"); err != nil {
		return err
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB, &models.Product{}, page))
}

func CreateProduct(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorize(c, "products"); err != nil {
		return err
	}

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
	if err := middlewares.IsAuthorize(c, "products"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	database.DB.Find(&product)

	return c.JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorize(c, "products"); err != nil {
		return err
	}

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
	if err := middlewares.IsAuthorize(c, "products"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: uint(id),
	}

	database.DB.Delete(&product)

	return c.JSON(nil)
}
