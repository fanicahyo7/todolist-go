package main

import (
	"database/sql"
	"log"
	"os"
	"todolist/handler"
	"todolist/repository"
	"todolist/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	//ambil data env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// connect to database
	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// create new fiber app
	app := fiber.New()

	// add middlewares
	app.Use(logger.New())
	app.Use(recover.New())

	todoListRepo := repository.NewTodoRepository(db)
	todoListSvc := service.NewTodoListService(todoListRepo)

	todohandler := handler.NewGroupHandler(todoListSvc)

	app.Get("/todos", todohandler.GetTodos)

	// start serverlanju
	log.Fatal(app.Listen(os.Getenv("SERVER_ADDRESS")))
}
