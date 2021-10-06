package api

import (
	"calcio/server/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Teams struct {
	log  *zap.SugaredLogger
	team *service.Team
}

// TeamsModule makes the injectable available for FX.
var TeamsModule = fx.Provide(NewTeams)

// NewTeams creates a new injectable.
func NewTeams(logger *zap.SugaredLogger, team *service.Team) *Teams {
	return &Teams{
		log:  logger,
		team: team,
	}
}

func (t Teams) Start(router fiber.Router, middlewares ...fiber.Handler) {
	for _, middleware := range middlewares {
		if middleware != nil {
			router.Use(middleware)
		}
	}

	router.Get("", t.list)
	router.Put("", t.create)
}

// @Summary List teams the users belongs to
// @Description Retrieve a list of teams as json, all teams if user is admin otherwise the teams the user belongs to
// @Tags teams
// @Accept json
// @Produce json
// @Success 200 {array} ent.Team "The list of teams"
// @Failure 400 {string} string "Authentication token is absent"
// @Failure 401 {string} string "Invalid authentication token"
// @Failure 500 {string} string "Something went wrong"
// @Param Authorization header string true "The authentication token"
// @Router /api/teams [get]
func (t Teams) list(ctx *fiber.Ctx) error {
	teams, err := t.team.List(ctx.UserContext())
	if err != nil {
		t.log.Error(err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(teams)
}

// @Summary Create a team
// @Description Create a new team with the connected user as the first member
// @Tags teams
// @Accept json
// @Produce json
// @Success 200 {object} ent.Team "The newly created team"
// @Failure 400 {string} string "Authentication token is absent"
// @Failure 401 {string} string "Invalid authentication token"
// @Failure 500 {string} string "Something went wrong, check that a team with the same name does not already exist"
// @Param Authorization header string true "The authentication token"
// @Param teammate query string false "The teammate id"
// @Router /api/teams [put]
func (t Teams) create(ctx *fiber.Ctx) error {
	teammate := ctx.Query("teammate")
	team, err := t.team.Create(teammate, ctx.UserContext())
	if err != nil {
		t.log.Error(err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(team)
}
