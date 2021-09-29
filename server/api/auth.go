package api

import (
	"time"

	"calcio/server/security"
	"calcio/server/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Auth struct {
	app      *fiber.App
	log      *zap.SugaredLogger
	uService *service.User
}

// AuthModule makes the injectable available for FX.
var AuthModule = fx.Provide(NewAuth)

// NewAuth creates a new injectable.
func NewAuth(app *fiber.App, logger *zap.SugaredLogger, user *service.User) *Auth {
	return &Auth{
		app:      app,
		log:      logger,
		uService: user,
	}
}

func (a Auth) Start(base string, middlewares ...fiber.Handler) {
	router := a.app.Group(base)
	for _, middleware := range middlewares {
		if middleware != nil {
			router.Use(middleware)
		}
	}

	router.Post("/login", a.login)
}

type login struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// @Summary Permits a user to log in to Calcio if credentials are valid.
// @Description Log in and retrieve the PASETO token signed. This method is rate limited.
// @Tags authentication
// @Accept json
// @Produce json
// @Success 200 {string} string "Paseto Token"
// @Failure 400 {string} string "Wrong login information provided"
// @Failure 429 {string} string "Rate limit reached"
// @Failure 500 {string} string "Something went wrong"
// @Param login body login true "Login json object"
// @Router /api/auth/login [post]
func (a Auth) login(ctx *fiber.Ctx) error {
	l := new(login)
	if err := ctx.BodyParser(l); err != nil {
		return fiber.ErrBadRequest
	}

	u, err := a.uService.Login(l.Name, l.Password)
	if err != nil {
		return fiber.ErrBadRequest
	}

	token, err := security.SignToken(security.Claims{
		UserId:  u.ID.String(),
		IsAdmin: u.Admin,
	}, 30*time.Minute)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.SendString(token)
}
