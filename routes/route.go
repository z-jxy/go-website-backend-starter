package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/z-jxy/blogbackend/controllers"
	"github.com/z-jxy/blogbackend/middleware"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("api/login", controllers.Login)

	// must be logged in to access follow routes
	app.Use(middleware.IsAuthenticated)

	app.Post("/api/post", controllers.CreatePost)
	app.Get("/api/allposts", controllers.AllPosts)
	app.Get("/api/post/:id", controllers.DetailPost)
	app.Put("/api/updatepost/:id", controllers.UpdatePost)
	app.Get("/api/user/posts", controllers.UserPosts)
	app.Delete("/api/user/deletepost/:id", controllers.DeletePost)
	app.Post("/api/upload-image", controllers.Upload)
	app.Static("/api/uploads", "./uploads")
}
