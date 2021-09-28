// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// GamesColumns holds the columns for the "games" table.
	GamesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "date", Type: field.TypeTime},
	}
	// GamesTable holds the schema information for the "games" table.
	GamesTable = &schema.Table{
		Name:       "games",
		Columns:    GamesColumns,
		PrimaryKey: []*schema.Column{GamesColumns[0]},
	}
	// ParticipationsColumns holds the columns for the "participations" table.
	ParticipationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "goals", Type: field.TypeInt, Default: 0},
		{Name: "participation_game", Type: field.TypeUUID, Nullable: true},
		{Name: "participation_team", Type: field.TypeUUID, Nullable: true},
	}
	// ParticipationsTable holds the schema information for the "participations" table.
	ParticipationsTable = &schema.Table{
		Name:       "participations",
		Columns:    ParticipationsColumns,
		PrimaryKey: []*schema.Column{ParticipationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "participations_games_game",
				Columns:    []*schema.Column{ParticipationsColumns[2]},
				RefColumns: []*schema.Column{GamesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "participations_teams_team",
				Columns:    []*schema.Column{ParticipationsColumns[3]},
				RefColumns: []*schema.Column{TeamsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TeamsColumns holds the columns for the "teams" table.
	TeamsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// TeamsTable holds the schema information for the "teams" table.
	TeamsTable = &schema.Table{
		Name:       "teams",
		Columns:    TeamsColumns,
		PrimaryKey: []*schema.Column{TeamsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "admin", Type: field.TypeBool, Default: false},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserTeamsColumns holds the columns for the "user_teams" table.
	UserTeamsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "team_id", Type: field.TypeUUID},
	}
	// UserTeamsTable holds the schema information for the "user_teams" table.
	UserTeamsTable = &schema.Table{
		Name:       "user_teams",
		Columns:    UserTeamsColumns,
		PrimaryKey: []*schema.Column{UserTeamsColumns[0], UserTeamsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_teams_user_id",
				Columns:    []*schema.Column{UserTeamsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_teams_team_id",
				Columns:    []*schema.Column{UserTeamsColumns[1]},
				RefColumns: []*schema.Column{TeamsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GamesTable,
		ParticipationsTable,
		TeamsTable,
		UsersTable,
		UserTeamsTable,
	}
)

func init() {
	ParticipationsTable.ForeignKeys[0].RefTable = GamesTable
	ParticipationsTable.ForeignKeys[1].RefTable = TeamsTable
	UserTeamsTable.ForeignKeys[0].RefTable = UsersTable
	UserTeamsTable.ForeignKeys[1].RefTable = TeamsTable
}
