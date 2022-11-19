package account

import (
	"github.com/jackc/pgx/v4"
)

func (api *API) ListAccounts(conn *pgx.Conn) ([]Account, error) {
	accounts := api.Store.ListAccounts(db)

	return accounts, nil
}

func (ag AccountGroup) GetAccount() error {
	return nil
}
