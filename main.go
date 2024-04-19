package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)



type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRoutes(app *fiber.App){
	api := app.Group("/api");

	api.POST("/create_books", r.CreateBook);
	api.GET("/books", r.GetBooks);
	api.GET("/get_book/:id", r.GetBook);
	api.DELETE("/delete_book/:id", r.DeleteBook);
}


func main(){

	err := godotenv.Load(".env");

	if err!= nil {
		log.Fatal(err);
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New();

	r.SetupRoutes(app)

	app.Listen(":8080")
}