package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/jbc2313/goRESTapi/db"
	"github.com/jbc2313/goRESTapi/models"
)

type User struct {
    ID uint `json:"id"`
    FirstName string `json:"first_name"`
    LastName string `json:"last_name"`
}

func CreateResponseUser(user models.User) User {
    return User{ID: user.ID, FirstName: user.FirstName, LastName: user.LastName}
}

func CreateUser(c *fiber.Ctx) error {
    var user models.User

    if err := c.BodyParser(&user); err != nil {
        return c.Status(400).JSON(err.Error())
    }

    db.Database.Db.Create(&user)
    responseUser := CreateResponseUser(user)
    return c.Status(200).JSON(responseUser)

}

func GetUsers(c *fiber.Ctx) error {
    users := []models.User{}
    db.Database.Db.Find(&users)
    responseUsers := []User{}
    for _, user := range users {
        responseUser := CreateResponseUser(user)
        responseUsers = append(responseUsers, responseUser)
    }

    return c.Status(200).JSON(responseUsers)
}

func findUser(id int, user *models.User) error {
    db.Database.Db.Find(&user, "id = ?", id)
    if user.ID == 0 {
        return errors.New("user doesn't exist")
    }
    return nil
}


func UpdateUser(c *fiber.Ctx) error {
    id, err := c.ParamsInt("id")

    var user models.User

    if err != nil {
        return c.Status(400).JSON("Please use an Integer")
    }

    err = findUser(id, &user)

    if err != nil {
        return c.Status(400).JSON(err.Error())
    }

    type UpdateUser struct {
        FirstName string `json:"first_name"`
        LastName string `json:"last_name"`
    }

    var updateData UpdateUser

    if err := c.BodyParser(&updateData); err != nil {
        return c.Status(500).JSON(err.Error())
    }

    user.FirstName = updateData.FirstName
    user.LastName = updateData.LastName

    db.Database.Db.Save(&user)

    responseUser := CreateResponseUser(user)

    return c.Status(200).JSON(responseUser)
}


func DeleteUser(c *fiber.Ctx) error {
    id, err := c.ParamsInt("id")

    var user models.User

    if err != nil {
        return c.Status(400).JSON("Please use an integer")
    }

    err = findUser(id, &user)

    if err != nil {
        return c.Status(400).JSON(err.Error())
    }

    if err = db.Database.Db.Delete(&user).Error; err != nil {
        return c.Status(404).JSON(err.Error())
    }
    return c.Status(200).JSON("User Deleted")
}
