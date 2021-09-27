package api

import (
	"calcio/ent"
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

	router.Get("", u.all)
}

// @Summary Fetch all Calcio's users.
// @Description Retrieves all Calcio's users as a json list.
// @Tags players
// @Accept json
// @Produce json
// @Success 200 {array} ent.User "The list of users"
// @Failure 400 {string} string "When the token is absent"
// @Failure 401 {string} string "When the token is invalid"
// @Failure 500 {string} string "When something went wrong"
// @Param Authorization header string true "The authentication token"
// @Router /api/users [get]
func (u Users) all(ctx *fiber.Ctx) error {
	var users = make([]ent.User, 0, 0)
	return ctx.JSON(users)
}
