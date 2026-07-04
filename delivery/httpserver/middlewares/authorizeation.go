package middlewares

import (
	"encoding/base64"
	"fmt"
	"os"
	"ostadbun/param/userparam"
	"ostadbun/service/userservice"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Auth(u userservice.User) func(c *fiber.Ctx) error {

	return func(c *fiber.Ctx) error {

		fmt.Println("mid start")
		token, IsBasicAuthMethod, err := GetAuthToken(c)

		fmt.Println(token, IsBasicAuthMethod)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "access denied",
			})
		}
		var userId int

		if IsBasicAuthMethod {

			//	we have email and password

			base64decoded, err := base64.StdEncoding.DecodeString(token)

			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "access denied",
				})
			}

			dta := strings.Split(string(base64decoded), ":")

			username := dta[0]
			password := dta[1]

			fmt.Println(username, password)
			ThisuserId, errEx := u.IsExist(userparam.User{
				Email:    username,
				Password: password,
			})

			fmt.Println(ThisuserId, errEx)
			if errEx != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "access denied",
				})
			}

			userId = int(*ThisuserId)

			fmt.Println("basic auth methud", username, password, errEx)

		} else {

			userID, err := u.AuthCheck(c.Context(), token)

			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "session not found",
					"error":   err.Error(),
				})
			}

			userId = userID
		}

		ID := strconv.Itoa(userId)
		c.Locals("user_id", ID)
		fmt.Println("1323", ID)

		fmt.Println("mid end")
		return c.Next()

	}

}

// the bool one 2th return value:
// Basic = true
// Brearer = false
func GetAuthToken(c *fiber.Ctx) (string, bool, error) {

	tkn := os.Getenv("COOKIE_TOKEN")

	cookieToken := c.Cookies(tkn)

	headerToken := c.Get("Authorization")

	if cookieToken != "" {
		return cookieToken, false, nil
	}

	if headerToken != "" {

		fmt.Println(headerToken, "header token")

		isBearer := strings.Contains(headerToken, "Bearer")

		isBasic := strings.Contains(headerToken, "Basic")

		if isBearer {
			headerToken = strings.Replace(headerToken, "Bearer ", "", 1)
			fmt.Println(headerToken, "is bearer")
			return headerToken, false, nil
		} else if isBasic {
			headerToken = strings.Replace(headerToken, "Basic ", "", 1)
			fmt.Println(headerToken, "is basic")
			return headerToken, true, nil
		}

	}

	return "", false, fmt.Errorf("user not authenticated")
}

type NewUseragentData struct {
	Id     int    `json:"id"`
	Type   string `json:"type"`
	Client string `json:"client"`
	Os     string `json:"os"`
}
