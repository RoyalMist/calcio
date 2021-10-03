package api

import (
	"calcio/ent"
	"calcio/server/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type user struct {
	ent.User
	Password string `json:"password"`
}

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
	router.Post("", u.create)
	router.Put("", u.update)
	router.Delete("/:id", u.delete)
}

// @Summary Fetch all Calcio's users.
// @Description Retrieves all Calcio's users as a json list.
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} ent.User "The list of users"
// @Failure 400 {string} string "Authentication token is absent"
// @Failure 401 {string} string "Invalid authentication token"
// @Failure 500 {string} string "Something went wrong"
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

// @Summary Create a new user.
// @Description Permits an administrator to create other users.
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} ent.User "The newly created user"
// @Failure 400 {string} string "Wrong input"
// @Failure 401 {string} string "Forbidden"
// @Failure 500 {string} string "Something went wrong"
// @Param Authorization header string true "The authentication token"
// @Param user body user true "The user to create"
// @Router /api/users [post]
func (u Users) create(ctx *fiber.Ctx) error {
	body := new(user)
	if err := ctx.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	usr, err := u.userService.Create(ent.User{
		Name:     body.Name,
		Password: body.Password,
		Admin:    body.Admin,
	}, ctx.UserContext())
	if err != nil {
		u.log.Error(err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(usr)
}

// @Summary Update an existing user.
// @Description Permits an administrator to update a user.
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} ent.User "The updated user"
// @Failure 400 {string} string "Wrong input"
// @Failure 401 {string} string "Forbidden"
// @Failure 500 {string} string "Something went wrong"
// @Param Authorization header string true "The authentication token"
// @Param user body user true "The user to update"
// @Router /api/users [put]
func (u Users) update(ctx *fiber.Ctx) error {
	body := new(user)
	if err := ctx.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	usr, err := u.userService.Update(ent.User{
		ID:       body.ID,
		Name:     body.Name,
		Password: body.Password,
		Admin:    body.Admin,
	}, ctx.UserContext())
	if err != nil {
		u.log.Error(err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(usr)
}

// @Summary Delete an existing user.
// @Description Permits an admin to delete a user.
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {string} string "The success of the operation"
// @Failure 400 {string} string "Wrong input"
// @Failure 401 {string} string "Forbidden"
// @Failure 500 {string} string "Something went wrong"
// @Param Authorization header string true "The authentication token"
// @Param id path string true "The id of the user to delete"
// @Router /api/users/{id} [delete]
func (u Users) delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	i, err := u.userService.Delete(id, ctx.UserContext())
	if err != nil {
		u.log.Error(err)
		return fiber.ErrInternalServerError
	}

	if i == 0 {
		return fiber.ErrBadRequest
	}

	return ctx.SendStatus(fiber.StatusOK)
}
