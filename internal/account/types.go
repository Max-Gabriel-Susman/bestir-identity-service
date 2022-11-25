package account

type Account struct {
	// ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	ID      int
	Balance int
}

type AccountGroup struct {
}
