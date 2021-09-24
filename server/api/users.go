package api

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Users struct {
	app *fiber.App
	log *zap.SugaredLogger
}

// UsersModule makes the injectable available for FX.
var UsersModule = fx.Provide(NewUsers)

// NewUsers creates a new injectable.
func NewUsers(app *fiber.App, logger *zap.SugaredLogger) *Users {
	return &Users{
		app: app,
		log: logger,
	}
}

func (u Users) Start(base string, middlewares ...fiber.Handler) {
	router := u.app.Group(base)
	for _, middleware := range middlewares {
		if middleware != nil {
			router.Use(middleware)
		}
	}

	router.Get("", u.All)
}

type dummyUsers struct {
	Name      string `json:"name"`
	IsAdmin   bool   `json:"isAdmin"`
	Mail      string `json:"mail"`
	AvatarUrl string `json:"avatarUrl"`
}

func (u Users) All(ctx *fiber.Ctx) error {
	users := []dummyUsers{
		{
			Name:      "Jane",
			IsAdmin:   true,
			Mail:      "jane@calcio.com",
			AvatarUrl: "https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=4&w=256&h=256&q=60",
		},
		{
			Name:      "John",
			IsAdmin:   false,
			Mail:      "john@calcio.com",
			AvatarUrl: "https://images.unsplash.com/photo-1570295999919-56ceb5ecca61?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=4&w=256&h=256&q=60",
		},
		{
			Name:      "Kristin",
			IsAdmin:   false,
			Mail:      "kristin@calcio.com",
			AvatarUrl: "https://images.unsplash.com/photo-1532417344469-368f9ae6d187?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=4&w=256&h=256&q=60",
		},
		{
			Name:      "Jonas",
			IsAdmin:   false,
			Mail:      "jonas@calcio.com",
			AvatarUrl: "https://images.unsplash.com/photo-1566492031773-4f4e44671857?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=4&w=256&h=256&q=60",
		},
	}

	return ctx.JSON(users)
}
