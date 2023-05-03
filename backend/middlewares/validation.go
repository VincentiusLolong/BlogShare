package middlewares

import (
	"mailinglist/infrastructure/security"

	"github.com/gofiber/fiber/v2"
)

// import (
// 	"regexp"

// 	"github.com/go-playground/validator/v10"
// )

// func CustomDateValidation(fl validator.FieldLevel) bool {
// 	dateStr := fl.Field().String()
// 	format := regexp.MustCompile(`^(\d{2})[-/](\d{2})[-/](\d{4})$`)
// 	return format.MatchString(dateStr)
// }

// func RegisterCustomValidators(v *validator.Validate) bool {
// 	err := v.RegisterValidation("birthdate", CustomDateValidation)
// 	if err != nil {
// 		return false
// 	}
// 	return true
// }

func Auth() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// c.Set("Cache-Control", "no-cache, no-store, must-revalidate")
		// c.Set("Pragma", "no-cache")
		// c.Set("Expires", "0")
		cookie := c.Cookies("jwt")
		if cookie == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"status":  "Unauthenticated",
				"message": "you're not login",
			})
		}
		data, claim := security.ValidateToken(cookie)
		if claim != nil {
			return c.Status(fiber.StatusForbidden).JSON(&fiber.Map{
				"status":  "ERROR",
				"message": claim.Error(),
			})
		}
		id := data["account_id"].(string)
		c.Locals("id", id)

		return c.Next()

	}
}
