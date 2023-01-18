package controller

import (
	"sully/todo-app/model"
	TodoRepository "sully/todo-app/todo-repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Controller struct {
	db *gorm.DB
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{db}
}

func (cont *Controller) GetAllTodoItems(c *fiber.Ctx) error {
	items, err := TodoRepository.FindAll(cont.db)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
	}
	return c.JSON(items)
}

func (cont *Controller) GetTodoItem(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		c.Status(fiber.StatusBadRequest).Send([]byte(err.Error()))
	}
	item, err := TodoRepository.FindById(cont.db, id)
	if err != nil {
		c.Status(fiber.StatusNotFound).Send([]byte(err.Error()))
	}
	return c.JSON(item)
}
func (cont *Controller) AddTodoItem(c *fiber.Ctx) error {
	todoItem := new(model.TodoItem)
	if err := c.BodyParser(todoItem); err != nil {
		c.Status(fiber.ErrBadRequest.Code).Send([]byte(err.Error()))
	}
	if err := TodoRepository.Save(cont.db, todoItem); err != nil {
		c.Status(fiber.StatusBadRequest).Send([]byte(err.Error()))
	}
	return c.JSON(todoItem)
}

func (cont *Controller) CompleteTodoItem(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Send([]byte(err.Error()))
	}
	if err := TodoRepository.Complete(cont.db, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
	}
	return c.SendStatus(fiber.StatusAccepted)
}

func (cont *Controller) DeleteTodoItem(c *fiber.Ctx) error {
	return nil
}
