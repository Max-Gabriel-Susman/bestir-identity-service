package account

import (
	"github.com/jackc/pgx/v4"
)

type API struct {
	// Logger *bestirlog.Logger
	// Store CockroachDBStorage // we'll do cockroach l8r
}

// we may want to parameterize storage and logging later
// func NewAPI(conn *pgx.Conn) *API {
func NewAPI(conn *pgx.Conn) *API {
	// return &API{Store: *NewCockroachDBStorage(conn)} // we'll do cockroach l8r
}
