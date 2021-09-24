package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"time"

	_ "calcio/docs"
	"calcio/ent"
	_ "calcio/ent/runtime"
	"calcio/server/api"
	"calcio/server/security"
	"calcio/server/settings"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"go.uber.org/fx"
)

//go:embed web/dist
var efs embed.FS

// @title Calcio API
// @version 1.0
// @description Calcio, Table Football App.
// @contact.name Royal Mist
// @contact.url https://github.com/RoyalMist
// @contact.email royalmist@calcio.ch
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
// @schemes http
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
	auth.Start("/api/auth", security.RateLimit(3, 10*time.Minute))
	users.Start("/api/users")
	app.Get("/swagger/*", swagger.Handler)

	web, err := fs.Sub(efs, "web/dist")
	if err != nil {
		log.Fatal(err)
	}

	app.Use(filesystem.New(filesystem.Config{
		Root:         http.FS(web),
		NotFoundFile: "index.html",
	}))
}
