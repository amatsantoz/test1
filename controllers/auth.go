package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"

	"achmad/shoppingcart/database"
	"achmad/shoppingcart/models"

	"golang.org/x/crypto/bcrypt"
)

// type LoginForm struct {
// 	Username     string  `form:"username" json:"username" validate:"required"`
// 	Password string  `form:"password" json:"password" validate:"required"`
// }

type AuthController struct {
	// declare variables
	Db *gorm.DB
	store *session.Store
}
func InitAuthController(s *session.Store) *AuthController {
	db := database.InitDb()
	// gorm
	db.AutoMigrate(&models.User{})

	return &AuthController{store: s, Db: db}
}

// get /register
func (controller *AuthController) Register(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{
		"Title": "Register Shopping Cart",
	})
}

// POST /user/create
func (controller *AuthController) PostRegister(c *fiber.Ctx) error {
	//myform := new(models.Product)
	
	var myform models.User
	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/register")
	}
	
	bytes, _ := bcrypt.GenerateFromPassword([]byte(myform.Password),10)
	sHash := string(bytes)
	formcreate := models.User {
		Name: myform.Name,
		Email: myform.Email,
		Username: myform.Username,
		Password: sHash,
	}
	// fmt.Println(sHash)
	// save product
	err := models.CreateUser(controller.Db, &formcreate)
	if err!=nil {
		return c.Redirect("/register")
	}
	// if succeed
	// fmt.Println(formcreate)
	return c.Redirect("/login")	
}

// get /login
func (controller *AuthController) Login(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "Login Shopping Cart",
	})
}

// post /login
func (controller *AuthController) PostLogin(c *fiber.Ctx) error {
	sess,err := controller.store.Get(c)
	if err!=nil {
		panic(err)
	}
	var myform models.User
	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/login")
	}

	// formcreate := models.User {
	// 	Username: myform.Username,
	// 	Password: sHash,
	// }

	
	// var formuname string = myform.Username
	// var formpass string = myform.Password
	var users models.User
	errr := models.ReadUserByUsername(controller.Db, &users, myform.Username )
	if errr!=nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	// hardcode auth
	if users.Username == myform.Username {
		compare := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(myform.Password))
		if compare == nil {
			sess.Set("username", myform.Username)
			sess.Save()
			fmt.Println("username dan pass cocok")
			return c.Redirect("/shopping/")
		} 
	}
	return c.Redirect("/login")
}

// /profile
func (controller *AuthController) Profile(c *fiber.Ctx) error {
	sess,err := controller.store.Get(c)
	if err!=nil {
		panic(err)
	}
	val := sess.Get("username")

	return c.JSON(fiber.Map{
		"username": val,
	})
}
// /logout
func (controller *AuthController) Logout(c *fiber.Ctx) error {
	
	sess,err := controller.store.Get(c)
	if err!=nil {
		panic(err)
	}
	sess.Destroy()
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}