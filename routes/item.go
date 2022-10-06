package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/jbc2313/goRESTapi/db"
	"github.com/jbc2313/goRESTapi/models"
)

type Item struct {
    ID uint `json:"id"`
    Name string `json:"name"`
}

func CreateResponseItem(item models.Item) Item {
    return Item{ID: item.ID, Name: item.Name}
}

func CreateItem(c *fiber.Ctx) error {
    var item models.Item

    if err := c.BodyParser(&item); err != nil {
        return c.Status(400).JSON(err.Error())
    }

    db.Database.Db.Create(&item)
    responeItem := CreateResponseItem(item)
    return c.Status(200).JSON(responeItem)
}

func GetItems(c *fiber.Ctx) error {
    items := []models.Item{}
    db.Database.Db.Find(&items)
    responseItems := []Item{}
    for _, item := range items {
        responseItem := CreateResponseItem(item)
        responseItems = append(responseItems, responseItem)
    }

    return c.Status(200).JSON(responseItems)

}

func findItem(id int, item *models.Item) error {
    db.Database.Db.Find(&item, "id = ?", id)
    if item.ID == 0 {
        return errors.New("item does not exist!")
    }
    return nil
}

func GetItem(c *fiber.Ctx) error {
    id, err := c.ParamsInt("id")

    var item models.Item

    if err != nil {
        return c.Status(400).JSON("Please enter an integer")
    }
    
    if err := findItem(id, &item); err != nil {
        return c.Status(400).JSON(err.Error())
    }

    responseItem := CreateResponseItem(item)

    return c.Status(200).JSON(responseItem)
}

func UpdateItem(c *fiber.Ctx) error {
    id, err := c.ParamsInt("id")

    var item models.Item

    if err != nil {
        return c.Status(400).JSON("Please inter an Integer")
    }

    err = findItem(id, &item)

    if err != nil {
        return c.Status(400).JSON(err.Error())
    }

    type UpdateItem struct {
        Name string `json:"name"`
    }

    var updateData UpdateItem

    if err := c.BodyParser(&updateData); err != nil {
        return c.Status(500).JSON(err.Error())
    }

    item.Name = updateData.Name

    db.Database.Db.Save(&item)

    responseItem := CreateResponseItem(item)

    return c.Status(200).JSON(responseItem)
}


