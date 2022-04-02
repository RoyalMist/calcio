package main

import (
	"time"

	_ "calcio/docs"
	_ "calcio/ent/runtime"
	"calcio/server/api"
	"calcio/server/security"
	"calcio/server/service"
	"calcio/server/settings"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/helmet/v2"
	"go.uber.org/fx"
)

// @title Calcio API
// @version 1.0
// @description Calcio, Table Football App
// @contact.name Royal Mist
// @contact.url https://github.com/RoyalMist
// @contact.email royalmist@calcio.ch
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @schemes http https
func main() {
	fx.New(
		fx.Options(
			settings.Module,
			api.AuthModule,
			api.UsersModule,
			api.TeamsModule,
			api.GamesModule,
			service.UserModule,
			service.TeamModule,
			service.GameModule,
		),
		fx.Invoke(setup),
		fx.Invoke(run),
	).Run()
}

func setup(user *service.User) {
	_ = user.CreateDefaultAdmin("admin123")
}

func run(app *fiber.App, auth *api.Auth, users *api.Users, teams *api.Teams) {
	apiRouter := app.Group("/api").Use(helmet.New())
	auth.Start(apiRouter.Group("/auth"), security.RateLimit(10, 5*time.Minute))
	users.Start(apiRouter.Group("/users"), security.IsAuthenticated)
	teams.Start(apiRouter.Group("/teams"), security.IsAuthenticated)
	apiRouter.Get("/doc/*", swagger.Handler)
}
