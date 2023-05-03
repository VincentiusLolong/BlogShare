package controllers

import (
	"fmt"
	"mailinglist/configs"
	"mailinglist/infrastructure/models"
	"mailinglist/infrastructure/security"
	"mailinglist/infrastructure/services"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Controllers interface {
	// user authentication
	SignUp(c *fiber.Ctx) error
	SignIn(c *fiber.Ctx) error
	Users(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error

	//user data
	CreateContent(c *fiber.Ctx) error
}

type controller struct {
	service services.Services
	c       fiber.Ctx
}

func New(service services.Services) Controllers {
	return &controller{service: service}
}

var (
	v            = validator.New()
	JwtSecretKey = configs.AllEnv("JWTKEY")
)

const (
	badr    = http.StatusBadRequest
	okr     = http.StatusOK
	unauth  = http.StatusUnauthorized
	serverr = http.StatusInternalServerError
)

func (ct *controller) status(status int, messages string, data any) error {
	return ct.c.Status(status).JSON(models.UserResponse{
		Status:  status,
		Message: messages,
		Data: &fiber.Map{
			"Result": data}})
}

func (ct *controller) SignUp(c *fiber.Ctx) error {
	var user models.User
	ct.c = *c

	if err := c.BodyParser(&user); err != nil {
		return ct.status(badr, "Json Input", err.Error())
	}

	if validationErr := v.Struct(&user); validationErr != nil {
		return ct.status(badr, "error", validationErr.Error())
	}

	pgxrow, errs := ct.service.SignUp(user)
	if errs != nil {
		return ct.status(badr, "error", errs.Error())
	}
	return ct.status(okr, "result", pgxrow)
}

func (ct *controller) SignIn(c *fiber.Ctx) error {
	var login models.Login
	ct.c = *c

	if err := c.BodyParser(&login); err != nil {
		return ct.status(badr, "Json Input", err.Error())
	}

	if validationErr := v.Struct(&login); validationErr != nil {
		return ct.status(badr, "error", validationErr.Error())
	}

	uid, errs := ct.service.Signin(login)
	if errs != nil {
		return ct.status(badr, "error", errs.Error())
	}
	claims, tokerr := security.GenerateJWT(models.GetDataToken{Account_id: uid})

	if tokerr != nil {
		fmt.Println(tokerr)
		switch tokerr.Error() {
		case "token contains an invalid number of segments":
			return ct.status(badr, "error", tokerr.Error())
		case "token is expired":
			return ct.status(unauth, "error", tokerr.Error())
		default:
			return ct.status(serverr, "error", tokerr.Error())
		}
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    claims,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return ct.status(okr, "Account Created", "200")
}

func (ct *controller) Users(c *fiber.Ctx) error {
	userid := fmt.Sprintf("%v", c.Locals("id"))
	ct.c = *c

	data, err := ct.service.HomepageUsers(userid)
	if err != nil {
		return ct.status(unauth, "error", err.Error())
	}

	return ct.status(okr, "success", data)
}

func (ct *controller) Logout(c *fiber.Ctx) error {
	ct.c = *c
	// err := ct.service.Logout()
	// if err != nil {
	// 	return ct.status(unauth, "error", err.Error())
	// }
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().AddDate(0, 0, -1),
		HTTPOnly: true,
		Path:     "/",
	})

	return ct.status(okr, "logout", "Success")
}
