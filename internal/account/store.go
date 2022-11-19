package account

import (
	"github.com/jackc/pgx/v4"
	"gorm.io/gorm"
)

type CockroachDBStorage struct {
	Conn *pgx.Conn
}

func NewCockroachDBStorage(conn *pgx.Conn) *CockroachDBStorage {
	return &CockroachDBStorage{Conn: conn}
}

func (c *CockroachDBStorage) ListAccounts(db *gorm.DB) []Account {
	var accounts []Account
	db.Find(&accounts)
	return accounts
}
