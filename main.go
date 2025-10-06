package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"todo-go/models"
	"todo-go/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Todo struct {
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) addTodo(context *fiber.Ctx) error {
	todo := Todo{}

	err := context.BodyParser(&todo)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Create(&todo).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create todo"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "todo has been added"})

	return nil
}

func (r *Repository) getTodos(context *fiber.Ctx) error {
	todoModels := &[]models.Todos{}

	err := r.DB.Find(todoModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the todos"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "todos fethcewd sucesfully",
		"data":    todoModels,
	})

	return nil
}

func (r *Repository) getTodoById(context *fiber.Ctx) error {
	id := context.Params("id")
	fmt.Println("The id is ", id)
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id is empty",
		})
		return nil
	}

	todoModel := &models.Todos{}

	err := r.DB.Where("id = ?", id).First(todoModel).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the todo"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book is fetched successfully",
		"data":    todoModel,
	})
	return nil
}

func (r *Repository) toggleTodoStatus(context *fiber.Ctx) error {
	id := context.Params("id")
	fmt.Println("The id is ", id)
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id is empty",
		})
		return nil
	}

	todoModel := &models.Todos{}

	err := r.DB.Where("id = ?", id).First(todoModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the todo"})
		return err
	}

	newStatus := !todoModel.Completed
	err2 := r.DB.Model(todoModel).Where("id = ?", id).Update("completed", newStatus).Error
	if err2 != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not update the todo"})
		return err2
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "todo status toggled successfully",
	})
	return nil

}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/add_todo", r.addTodo)
	api.Get("/toggle_todo/:id", r.toggleTodoStatus)
	api.Get("/get_todo", r.getTodos)
	api.Get("/get_todo/:id", r.getTodoById)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("could not load the database.")
	}

	err = models.MigrateTodos(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")

}
