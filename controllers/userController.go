package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/javiergomezve/backend-admin/database"
	"github.com/javiergomezve/backend-admin/middlewares"
	"github.com/javiergomezve/backend-admin/models"
)

func AllUsers(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorize(c, "users"); err != nil {
		return err
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB, &models.User{}, page))
}

func CreateUser(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorize(c, "users"); err != nil {
		return err
	}

	var userDTO fiber.Map

	if err := c.BodyParser(&userDTO); err != nil {
		return err
	}

	list := userDTO["roles"].([]interface{})
	roles := make([]models.Role, len(list))

	for i, roleId := range list {
		id, _ := strconv.Atoi(roleId.(string))

		roles[i] = models.Role{
			Id: uint(id),
		}
	}

	user := models.User{
		FirstName: userDTO["first_name"].(string),
		LastName:  userDTO["last_name"].(string),
		Email:     userDTO["email"].(string),
		Roles:     roles,
	}
	user.SetPassword("1234")

	database.DB.Create(&user)

	return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorize(c, "users"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id:    uint(id),
		Roles: []models.Role{},
	}

	database.DB.Preload("Role").Find(&user)

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorize(c, "users"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	var userDTO fiber.Map

	if err := c.BodyParser(&userDTO); err != nil {
		return err
	}

	list := userDTO["roles"].([]interface{})
	roles := make([]models.Role, len(list))

	for i, roleId := range list {
		id, _ := strconv.Atoi(roleId.(string))

		roles[i] = models.Role{
			Id: uint(id),
		}
	}

	result := make(map[string]string)
	database.DB.Table("role_user").Where("user_id", id).Delete(&result)

	user := models.User{
		Id:        uint(id),
		FirstName: userDTO["first_name"].(string),
		LastName:  userDTO["last_name"].(string),
		Email:     userDTO["email"].(string),
		Roles:     roles,
	}

	database.DB.Debug().Model(&user).Updates(user)

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	if err := middlewares.IsAuthorize(c, "users"); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB.Delete(&user)

	return nil
}
