package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html"

	"achmad/shoppingcart/controllers"
)

func main() {
	// session
	store := session.New()

	// load template engine
	engine := html.New("./views",".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// static
	app.Static("/public","./public")

	// controllers
	prodController := controllers.InitProductController(store)
	authController := controllers.InitAuthController(store)

	// p := app.Group("/greetings")
	// p.Get("/", helloController.Greeting)
	// p.Get("/hello", helloController.SayHello)
	// p.Get("/myview", helloController.HelloView)

	prod := app.Group("/shopping")
	prod.Get("/", prodController.Home)
	prod.Get("/create", prodController.AddProduct)
	prod.Post("/create", prodController.AddPostedProduct)
	prod.Get("/productdetail", prodController.GetDetailProduct)
	prod.Get("/detail/:id", prodController.GetDetailProduct2)
	prod.Get("/editproduct/:id", prodController.EditlProduct)
	prod.Post("/editproduct/:id", prodController.EditlPostedProduct)
	prod.Get("/deleteproduct/:id", prodController.DeleteProduct)

	app.Get("/register",authController.Register)
	app.Post("/register",authController.PostRegister)
	app.Get("/login",authController.Login)
	app.Post("/login",authController.PostLogin)
	app.Get("/logout",authController.Logout)
	app.Get("/profile",authController.Profile)

	// app.Use("/profile", func(c *fiber.Ctx) error {
	// 	sess,_ := store.Get(c)
	// 	val := sess.Get("username")
	// 	if val != nil {
	// 		return c.Next()
	// 	}

	// 	return c.Redirect("/login")

	// })
	app.Get("/profile", func(c *fiber.Ctx) error {
		sess,_ := store.Get(c)
		val := sess.Get("username")
		if val != nil {
			return c.Next()
		}

		return c.Redirect("/login")

	}, authController.Profile)

	app.Listen(":3000")
}