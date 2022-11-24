package database

import (
	"context"
	"database/sql/driver"
)

type connector struct {
	cfg *Config // immutable private copy.
}

// Connect implements driver.Connector interface.
// Connect returns a connection to the database.
func (c *connector) Connect(ctx context.Context) (driver.Conn, error) {
	pgc := &postgresConn{}
	return pgc, nil
}

func (pgc *postgresConn) Begin() (driver.Tx, error) {
	return pgc.begin(false)
}

func (pgc *postgresConn) begin(readOnly bool) (driver.Tx, error) {
	/*
		if pgc.closed.Load() {
			errLog.Print(ErrInvalidConn)
			return nil, driver.ErrBadConn
		}
		var q string
		if readOnly {
			q = "START TRANSACTION READ ONLY"
		} else {
			q = "START TRANSACTION"
		}
		err := mc.exec(q)
		if err == nil {
			return &postgresTx{pgc}, err
		}
		return nil, mc.markBadConn(err)
	*/
	if readOnly {
		return &postgresTx{pgc}, nil
	} else {
		return &postgresTx{pgc}, nil
	}
}

func (pgc *postgresConn) Prepare(query string) (driver.Stmt, error) {
	/*
		if mc.closed.Load() {
			errLog.Print(ErrInvalidConn)
			return nil, driver.ErrBadConn
		}
		// Send command
		err := mc.writeCommandPacketStr(comStmtPrepare, query)
		if err != nil {
			// STMT_PREPARE is safe to retry.  So we can return ErrBadConn here.
			errLog.Print(err)
			return nil, driver.ErrBadConn
		}
	*/
	stmt := &postgreStmt{
		pgc: pgc,
	}
	/*
		// Read Result
		columnCount, err := stmt.readPrepareResultPacket()
		if err == nil {
			if stmt.paramCount > 0 {
				if err = mc.readUntilEOF(); err != nil {
					return nil, err
				}
			}

			if columnCount > 0 {
				err = mc.readUntilEOF()
			}
		}
	*/
	return stmt, err
}

// Driver implements driver.Connector interface.
// Driver returns &MySQLDriver{}.
func (c *connector) Driver() driver.Driver {
	return &PostgresDriver{}
}
