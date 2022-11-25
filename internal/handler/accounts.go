package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Max-Gabriel-Susman/bestir-identity-service/internal/account"
	"github.com/Max-Gabriel-Susman/bestir-identity-service/internal/foundation/web"
	"github.com/go-chi/chi/v5"
)

type Account struct {
	ID string `json:"id"`
}

type accountGroup struct {
	*account.API
}

type ListAccountsResponse struct {
	Accounts []account.Account `json:"accounts"`
}

func AccountEndpoints(app *web.App, api *account.API) {
	ag := accountGroup{API: api}

	// app.Handle("GET", "/accounts", ag.ListAccounts)
	app.Handle("GET", "/", ag.ListAccounts)
}

func (ag accountGroup) ListAccounts(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	fmt.Println("ListAccounts invoked")
	// accounts, err := ag.API.ListAccounts(ag.API.Store.Conn)
	// if err != nil {
	// 	return err
	// }
	accounts := []account.Account{
		{
			ID:      69,
			Balance: 69,
		},
		{
			ID:      420,
			Balance: 420,
		},
	}

	fmt.Println("you got dem accounts fella")
	return web.Respond(ctx, w, ListAccountsResponse{
		Accounts: accounts,
	}, http.StatusOK)
}

func (ag accountGroup) GetAccount(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	accountID := chi.URLParam(r, "account_id")
	if accountID == "" {
		return nil
		// return handleMissingURLParameter(ctx, accountID, Account)
	}

	return nil
}
