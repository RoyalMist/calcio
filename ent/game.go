// Code generated by entc, DO NOT EDIT.

package ent

import (
	"calcio/ent/game"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Game is the model entity for the Game schema.
type Game struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Date holds the value of the "date" field.
	Date time.Time `json:"date,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GameQuery when eager-loading is set.
	Edges GameEdges `json:"edges"`
}

// GameEdges holds the relations/edges for other nodes in the graph.
type GameEdges struct {
	// Participations holds the value of the participations edge.
	Participations []*Participation `json:"participations,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ParticipationsOrErr returns the Participations value or an error if the edge
// was not loaded in eager-loading.
func (e GameEdges) ParticipationsOrErr() ([]*Participation, error) {
	if e.loadedTypes[0] {
		return e.Participations, nil
	}
	return nil, &NotLoadedError{edge: "participations"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Game) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case game.FieldDate:
			values[i] = new(sql.NullTime)
		case game.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Game", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Game fields.
func (ga *Game) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case game.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ga.ID = *value
			}
		case game.FieldDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date", values[i])
			} else if value.Valid {
				ga.Date = value.Time
			}
		}
	}
	return nil
}

// QueryParticipations queries the "participations" edge of the Game entity.
func (ga *Game) QueryParticipations() *ParticipationQuery {
	return (&GameClient{config: ga.config}).QueryParticipations(ga)
}

// Update returns a builder for updating this Game.
// Note that you need to call Game.Unwrap() before calling this method if this Game
// was returned from a transaction, and the transaction was committed or rolled back.
func (ga *Game) Update() *GameUpdateOne {
	return (&GameClient{config: ga.config}).UpdateOne(ga)
}

// Unwrap unwraps the Game entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ga *Game) Unwrap() *Game {
	tx, ok := ga.config.driver.(*txDriver)
	if !ok {
		panic("ent: Game is not a transactional entity")
	}
	ga.config.driver = tx.drv
	return ga
}

// String implements the fmt.Stringer.
func (ga *Game) String() string {
	var builder strings.Builder
	builder.WriteString("Game(")
	builder.WriteString(fmt.Sprintf("id=%v", ga.ID))
	builder.WriteString(", date=")
	builder.WriteString(ga.Date.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Games is a parsable slice of Game.
type Games []*Game

func (ga Games) config(cfg config) {
	for _i := range ga {
		ga[_i].config = cfg
	}
}
