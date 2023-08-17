package routes

import (
	"github.com/zatrano/zatrano/internal/app/handlers"

	"github.com/gofiber/fiber/v2"
)

type UserRouter struct{}

func (userRouter *UserRouter) SetupRoutes(router fiber.Router) {
	books := router.Group("/books")
	books.Get("/", handlers.GetAllBooks)
	books.Get("/:id", handlers.GetBookByID)
	books.Post("/", handlers.CreateBook)
	books.Put("/:id", handlers.UpdateBook)
	books.Delete("/:id", handlers.DeleteBook)
}