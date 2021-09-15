package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"calcio/api/settings"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

//go:embed web/dist
var efs embed.FS

func main() {
	fx.New(
		fx.Options(settings.Module),
		fx.Invoke(run),
	).Run()
}

func run(logger *zap.SugaredLogger, app *fiber.App) {
	logger.Info("Hello Audacia !")
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
