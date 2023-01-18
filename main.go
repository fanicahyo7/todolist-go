package main

import (
	"database/sql"
	"fmt"
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
	user, err := todoListRepo.GetByUserID(15)
	if err != nil {
		fmt.Println(err)
	} else {
		for a := 0; a <= len(user)-1; a++ {
			fmt.Println(user)
		}
	}

	userRepo := repository.NewUserRepository(db)

	// service
	todoListService := service.NewTodoListService(todoListRepo)
	userService := service.NewUserService(userRepo)

	// handler
	todohandler := handler.NewGroupHandler(todoListService)
	userHandler := handler.NewuserHandler(userService)

	// route
	authGroup := app.Group("/auth", util.JWTAuthMiddleware(os.Getenv("JWT_SECRET")))
	authGroup.Get("/todos", todohandler.GetTodos)
	app.Post("/register", userHandler.RegisterUser)

	// start serverlanju
	log.Fatal(app.Listen(os.Getenv("SERVER_ADDRESS")))
}
