package middlewares

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/javiergomezve/backend-admin/database"
	"github.com/javiergomezve/backend-admin/models"
	"github.com/javiergomezve/backend-admin/util"
)

func IsAuthorize(c *fiber.Ctx, page string) error {
	hasAccess := false

	id := util.ParseJwt(c)
	user := models.User{
		Id: uint(id),
	}
	database.DB.Preload("Roles").Preload("Permissions").Find(&user)

	if c.Method() == "GET" {
		var permissions []models.Permission
		database.DB.Preload("Roles").Where("name IN ?", []string{"view_" + page, "edit_" + page}).Find(&permissions)

		for _, permission := range permissions {
			fmt.Println("Permission required: ", permission.Name)
			for _, role := range permission.Roles {
				if hasAccess {
					break
				}

				for _, userRole := range user.Roles {
					fmt.Println("User role: ", userRole.Id)
					if role.Id == userRole.Id {
						hasAccess = true
						break
					}
				}
			}
		}
	} else {
		var permissions []models.Permission
		database.DB.Preload("Roles").Where("name IN ?", []string{"edit_" + page}).Find(&permissions)

		for _, permission := range permissions {
			for _, role := range permission.Roles {
				if hasAccess {
					break
				}

				for _, userRole := range user.Roles {
					if role.Id == userRole.Id {
						hasAccess = true
						break
					}
				}
			}
		}
	}

	if hasAccess {
		return nil
	} else {
		c.Status(fiber.StatusUnauthorized)
		return errors.New("Unauthorized")
	}
}
