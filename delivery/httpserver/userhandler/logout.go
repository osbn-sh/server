package userhandler

import (
	"fmt"
	"os"
	notify "ostadbun/pkg/bale/notif"
	"ostadbun/pkg/enviroment"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) logout(c *fiber.Ctx) error {

	//TODO remove user session

	htonlyKey := os.Getenv("COOKIE_TOKEN")
	usrnameKey := os.Getenv("COOKIE_NAME")

	cookeRemover(c, htonlyKey, usrnameKey)

	return c.Redirect(os.Getenv("WEB_CLIENT"))
}

func cookeRemover(c *fiber.Ctx, key1, key2 string) {

	fmt.Println("trying to remove cookie", key1, key2)
	c.Cookie(&fiber.Cookie{
		Name:     key1,
		Value:    "",
		Path:     "/",
		Domain:   os.Getenv("COOKIE_DOMAIN"),
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: enviroment.IsProduction(),
		//TODO make true on production https
		Secure:   enviroment.IsProduction(),
		SameSite: fiber.CookieSameSiteNoneMode,
	})

	c.Cookie(&fiber.Cookie{
		Name:    key2,
		Value:   "",
		Domain:  os.Getenv("COOKIE_DOMAIN"),
		Path:    "/",
		Expires: time.Now().Add(-time.Hour),
	})

	errnotif := notify.Notify(fmt.Sprintf("logout %s", key1, key2))

	fmt.Println(errnotif)

}
