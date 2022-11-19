package account

import "github.com/google/uuid"

type Account struct {
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Balance int
}

type AccountGroup struct {
}
