package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	_ "calcio/ent/runtime"
	"calcio/server/api"
	"calcio/server/settings"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"go.uber.org/fx"
)

//go:embed web/dist
var efs embed.FS

func main() {
	fx.New(
		fx.Options(
			settings.Module,
			api.Module,
		),
		fx.Invoke(run),
	).Run()
}

func run(app *fiber.App) {
	web, err := fs.Sub(efs, "web/dist")
	if err != nil {
		log.Fatal(err)
	}

	app.Get("/api/hello", sayHello)

	app.Use(filesystem.New(filesystem.Config{
		Root:         http.FS(web),
		NotFoundFile: "index.html",
	}))
}

type hello struct {
	Message string `json:"message"`
}

func sayHello(c *fiber.Ctx) error {
	h := hello{Message: "Hello Audacia"}
	return c.JSON(h)
}
