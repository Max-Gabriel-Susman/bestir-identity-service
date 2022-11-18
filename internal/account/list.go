package account

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (ag AccountGroup) ListAccounts(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	accountID := chi.URLParam(r, "account_id")
	if accountID == "" {
		return nil
		// return handleMissingURLParameter(ctx, accountID, Account)
	}

	// accounts, err := ag.GetAccounts()

	fmt.Println("you got dem accounts fella")
	return nil
}

func (ag AccountGroup) GetAccount() error {
	return nil
}
