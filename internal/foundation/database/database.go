package database

import (
	// "database/sql"
	// sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
	// "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	// "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	// postgres "github.com/jackc/pgx/v4"
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"

	dbr "github.com/gocraft/dbr/v2"
	"github.com/gocraft/dbr/v2/dialect"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
)

// Config is used to connect to db
type Config struct {
	User     string
	Password string
	Host     string
	Name     string
	Params   string
}

// Opens a database connection with configuration
func Open(cfg Config, serviceName string) (*sql.DB, error) {
	sqltrace.Register("postgres", &PostgresDriver{},
		sqltrace.WithServiceName(serviceName),
		sqltrace.WithAnalytics(true),
	)
	dsn := DSN(cfg)
	return sqltrace.Open("postgres", dsn)
}

type PostgresDriver struct{}

// func (pgd PostgresDriver) Open(name string) (PostgresConnection, error) {
func (pgd PostgresDriver) Open(dsn string) (driver.Conn, error) {
	cfg, err := ParseDSN(dsn)
	if err != nil {
		return nil, err
	}
	c := &connector{
		cfg: cfg,
	}
	return c.Connect(context.Background())
}

func NewDBR(db *sql.DB) *dbr.Connection {
	return &dbr.Connection{DB: db, EventReceiver: &dbr.NullEventReceiver{}, Dialect: dialect.PostgreSQL}
}

// Data Source Name - used to request a connection to a data source
func DSN(cfg Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", cfg.User, cfg.Password, cfg.Host, cfg.Name, cfg.Params)
}

// TRASH HEAP

/*

type PostgresConnection struct {
}


func (pgc PostgresConnection) Prepare(query string) (Stmt, error) {

	return Stmt{}, nil
}

func (pgc PostgresConnection) Close() error {
	return nil
}

func (pgc PostgresConnection) Begin() (PostgresTransaction, error) {
	return PostgresTransaction{}, nil
}

type PostgresTransaction struct {
}

func (pgt PostgresTransaction) Commit() error {
	return nil
}

func (pgt PostgresTransaction) Rollback() error {
	return nil
}

*/
