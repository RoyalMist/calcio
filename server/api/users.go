package api

import (
	"calcio/server/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Users struct {
	app         *fiber.App
	log         *zap.SugaredLogger
	userService *service.User
}

// UsersModule makes the injectable available for FX.
var UsersModule = fx.Provide(NewUsers)

// NewUsers creates a new injectable.
func NewUsers(app *fiber.App, logger *zap.SugaredLogger, user *service.User) *Users {
	return &Users{
		app:         app,
		log:         logger,
		userService: user,
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
	router.Post("/create", u.create)
	router.Put("create", u.update)
	router.Delete("/:id", u.delete)
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
	users, err := u.userService.List(ctx.UserContext())
	if err != nil {
		u.log.Errorf("impossible to retieve users %v", err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(users)
}

func (u Users) create(ctx *fiber.Ctx) error {
	return ctx.SendString("")
}

func (u Users) update(ctx *fiber.Ctx) error {
	return ctx.SendString("")
}

func (u Users) delete(ctx *fiber.Ctx) error {
	return ctx.SendString("")
}
