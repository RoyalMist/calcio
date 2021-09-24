package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"time"

	"calcio/ent"
	_ "calcio/ent/runtime"
	"calcio/server/api"
	"calcio/server/security"
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
			api.UsersModule,
		),
		fx.Invoke(setup),
		fx.Invoke(run),
	).Run()
}

func setup(client *ent.Client) {
	_ = security.CreateAdmin(client)
}

func run(app *fiber.App, auth *api.Auth, users *api.Users) {
	web, err := fs.Sub(efs, "web/dist")
	if err != nil {
		log.Fatal(err)
	}

	auth.Start("/auth", security.RateLimit(5, 2*time.Minute))
	users.Start("/users")
	app.Use(filesystem.New(filesystem.Config{
		Root:         http.FS(web),
		NotFoundFile: "index.html",
	}))
}
