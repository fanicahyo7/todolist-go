package main

import (
	"database/sql"
	"log"
	"os"
	"todolist/handler"
	"todolist/repository"
	"todolist/service"
	"todolist/util"

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

	// repo
	todoListRepo := repository.NewTodoRepository(db)
	userRepo := repository.NewUserRepository(db)

	// service
	todoListService := service.NewTodoListService(todoListRepo)
	userService := service.NewUserService(userRepo)

	// handler
	todohandler := handler.NewGroupHandler(todoListService)
	userHandler := handler.NewuserHandler(userService)

	// route
	// authGroup := app.Group("/auth", util.JWTAuthMiddleware(os.Getenv("JWT_SECRET")))
	// authGroup.Get("/todos", todohandler.GetTodos)
	app.Get("/todos", util.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), todohandler.GetTodos)
	app.Post("/register", userHandler.RegisterUser)
	app.Post("/login", userHandler.LoginUser)

	// start serverlanju
	log.Fatal(app.Listen(os.Getenv("SERVER_ADDRESS")))
}
