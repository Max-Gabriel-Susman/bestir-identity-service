package handler

import (
	"github.com/jackc/pgx/v4"
)

type Deps struct {
	// Logger // must have do eet
	Conn *pgx.Conn
}
