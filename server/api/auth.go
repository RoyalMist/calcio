package api

import (
	"time"

	"calcio/ent"
	"calcio/ent/user"
	"calcio/server/security"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Auth struct {
	app    *fiber.App
	log    *zap.SugaredLogger
	client *ent.Client
}

// AuthModule makes the injectable available for FX.
var AuthModule = fx.Provide(NewAuth)

// NewAuth creates a new injectable.
func NewAuth(app *fiber.App, logger *zap.SugaredLogger, client *ent.Client) *Auth {
	return &Auth{
		app:    app,
		log:    logger,
		client: client,
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
// @Failure 400 {string} string "When the token is absent or malformed"
// @Failure 401 {string} string "When the token is invalid"
// @Failure 429 {string} string "When the rate limit is reached"
// @Failure 500 {string} string "When something went wrong"
// @Param login body login true "Login json object"
// @Router /api/auth/login [post]
func (a Auth) login(ctx *fiber.Ctx) error {
	l := new(login)
	if err := ctx.BodyParser(l); err != nil {
		return fiber.ErrBadRequest
	}

	u, err := a.client.User.Query().Where(user.Name(l.Name)).First(ctx.UserContext())
	if err != nil {
		return fiber.ErrBadRequest
	}

	if !security.CheckPassword(l.Password, u.Password) {
		return fiber.ErrBadRequest
	}

	token, err := security.SignToken(security.Claims{
		UserId:  u.ID.String(),
		IsAdmin: u.Admin,
	}, 20*time.Minute)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.SendString(token)
}
