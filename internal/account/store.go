package account

import (
	"context"
	"fmt"

	"github.com/Max-Gabriel-Susman/bestir-identity-service/internal/foundation/database"
	"github.com/gocraft/dbr/v2"
)

// "github.com/jackc/pgx/v4" // cockroach my love, we will meet again

/*
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
*/

func NewMySQLStore(conn *dbr.Connection) *MySQLStorage {
	return &MySQLStorage{conn: conn, sess: conn.NewSession(nil)}
}

type MySQLStorage struct {
	conn *dbr.Connection
	sess *dbr.Session
}

var (
	accountTable = database.NewTable("account", Account{})
)

func (s *MySQLStorage) ListAccounts(ctx context.Context) ([]Account, error) {
	query := s.sess.Select(accountTable.Columns...).
		From(accountTable.Name)

	accounts := []Account{}

	if _, err := query.LoadContext(ctx, &accounts); err != nil {
		return accounts, database.ClassifyError(err)
	}

	return accounts, nil
}

func (s *MySQLStorage) getAccountByIdempotencyKey(ctx context.Context, idempotencyKey string) (Account, error) {
	var account Account
	err := s.sess.Select(accountTable.Columns...).
		From(accountTable.Name).
		Where("idempotency_key = ?", idempotencyKey).
		LoadOneContext(ctx, &account)
	return account, database.ClassifyError(err)
}

func (s *MySQLStorage) CreateAccount(ctx context.Context, account Account) error {
	fmt.Println("Create Account C invoked")
	_, err := s.sess.InsertInto(accountTable.Name).
		Columns(accountTable.Columns...).
		Record(account).
		ExecContext(ctx)
	return database.ClassifyError(err)
}
