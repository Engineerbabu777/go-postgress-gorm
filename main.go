package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)


type Book struct {
	Author string    `json:"author"`
	Title string    `json:"title"`
	Publisher string    `json:"publisher"`
}

type Repository struct {
	DB *gorm.DB
}

func(r Repository) CreateBook(context *fiber.Ctx) error{
	book := Book{};

	err := context.BodyParser(&book);

	if err!= nil {
		context.Status(400).JSON(&context.Map{"message":"request failed!"})
		return err;
	}

	err = r.DB.Create(&book).Error;

	if err!= nil {
		context.Status(404).JSON(&fiber.Map{"message":"could not create book"})
		return err;
	}

	context.Status(200).JSON(&fiber.Map{"message":"book added success"});

	return nil;
}

func (r *Repository) GetBooks(context *fiber.Ctx) error{
    var books []Book;

	err := r.DB.Find(&books).Error;

	if err!= nil {
		context.Status(404).JSON(&fiber.Map{"message":"could not get books"})
		return err;
	}

	context.Status(200).JSON(&books);
	return nil;
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

	db, err := storage.NewConnection(config)

	if err!= nil {
		log.Fatal("Could not create a connection")
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New();

	r.SetupRoutes(app)

	app.Listen(":8080")
}