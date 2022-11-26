package account

import (
	"github.com/google/uuid"
)

type Account struct {
	ID   uuid.UUID `db:"id" json:"id"`
	Name string    `db:"name" json:"name"`
}

type IncomingAccount struct {
	Name string `json:"name" required:"true"`
	// IdempotencyKey null.String `json:"-" db:"idempotency_key"`
}

type AccountGroup struct {
}
