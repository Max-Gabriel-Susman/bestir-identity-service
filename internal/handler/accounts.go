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
	app.Handle("GET", "/accounts", ag.ListAccounts)
	app.Handle("POST", "/account", ag.CreateAccount)
}

func (ag accountGroup) ListAccounts(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	accounts, err := ag.API.ListAccounts(ctx)
	if err != nil {
		return err
	}

	return web.Respond(ctx, w, ListAccountsResponse{
		Accounts: accounts,
	}, http.StatusOK)
}

// accounts := []account.Account{
// 	{
// 		ID:      69,
// 		Balance: 69,
// 	},
// 	{
// 		ID:      420,
// 		Balance: 420,
// 	},
// }

func (ag accountGroup) CreateAccount(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Create Account a invoked")
	var input account.IncomingAccount
	if err := web.Decode(r.Body, &input); err != nil {
		return err
	}

	account, err := ag.API.CreateAccount(ctx, input)
	if err != nil {
		return err
	}

	return web.Respond(ctx, w, account, http.StatusCreated)
}

func (ag accountGroup) GetAccount(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	accountID := chi.URLParam(r, "account_id")
	if accountID == "" {
		return nil
		// return handleMissingURLParameter(ctx, accountID, Account)
	}

	return nil
}
