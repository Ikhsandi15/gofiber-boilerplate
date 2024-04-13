package controllers

import "github.com/gofiber/fiber/v2"

func CreateUser(c *fiber.Ctx) error {
	return nil
}

func UserList(c *fiber.Ctx) error {
	return c.SendString("User list")
}

func GetUserDetails(c *fiber.Ctx) error {
	return nil
}

func EditUser(c *fiber.Ctx) error {
	return nil
}

func UpdateUser(c *fiber.Ctx) error {
	return nil
}

func DeleteUser(c *fiber.Ctx) error {
	return nil
}
