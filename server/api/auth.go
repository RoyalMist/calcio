package api

import (
	"time"

	"calcio/server/security"
	"calcio/server/service"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	log     *zap.SugaredLogger
	service *service.User
}

// AuthModule makes the injectable available for FX.
var AuthModule = fx.Provide(NewAuth)

// NewAuth creates a new injectable.
func NewAuth(logger *zap.SugaredLogger, user *service.User) *Auth {
	return &Auth{
		log:     logger,
		service: user,
	}
}

func (a Auth) Start(router fiber.Router, middlewares ...fiber.Handler) {
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
	body := new(login)
	if err := ctx.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	usr, err := a.service.Login(body.Name, body.Password)
	if err != nil {
		// simulate a bcrypt round to eliminate timing guesses.
		_, _ = bcrypt.GenerateFromPassword([]byte(uuid.NewString()), security.HashCost)
		return fiber.ErrBadRequest
	}

	token, err := security.SignToken(security.Claims{
		UserId:   usr.ID.String(),
		UserName: usr.Name,
		IsAdmin:  usr.Admin,
	}, 30*time.Minute)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return ctx.SendString(token)
}
