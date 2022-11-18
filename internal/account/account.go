package account

import (
	"context"
	"fmt"
	"net/http"
)

type Account struct {
	ID string `json:"id"`
}

type AccountGroup struct {
}

func (ag AccountGroup) GetAccounts(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	fmt.Println("you got dem accounts fella")
	return nil
}

func (ag AccountGroup) GetAccount() error {
	return nil
}
