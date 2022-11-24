package account

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

func (api *API) ListAccounts(conn *pgx.Conn) ([]Account, error) {
	rows, err := conn.Query(context.Background(), "SELECT id, balance FROM accounts")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id uuid.UUID
		var balance int
		if err := rows.Scan(&id, &balance); err != nil {
			log.Fatal(err)
		}
		log.Printf("%s: %d\n", id, balance)
	}
	// return nil

	return []Account{}, nil
}

func (ag AccountGroup) GetAccount() error {
	return nil
}
