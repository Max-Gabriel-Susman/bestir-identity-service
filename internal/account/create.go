package account

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (api *API) CreateAccount(ctx context.Context, incomingAccount IncomingAccount) (Account, error) {
	fmt.Println("Create Account B invoked")
	id := uuid.New()

	account := Account{
		ID:   id,
		Name: incomingAccount.Name,
	}

	err := api.Store.CreateAccount(ctx, account)

	// if bestirerror.StatusCode(err) == http.StatusConflict && incomingAccount.IdempotencyKey.String != "" {
	// 	fmt.Println("Create Account B failed")
	// 	return api.Store.getAccountByIdempotencyKey(ctx, incomingAccount.IdempotencyKey.String)
	// }
	return account, err
}
