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
			api.AuthModule,
		),
		fx.Invoke(run),
	).Run()
}

func run(app *fiber.App, auth *api.Auth) {
	web, err := fs.Sub(efs, "web/dist")
	if err != nil {
		log.Fatal(err)
	}

	auth.Start("/auth")

	app.Use(filesystem.New(filesystem.Config{
		Root:         http.FS(web),
		NotFoundFile: "index.html",
	}))
}
