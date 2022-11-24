package database

type postgreStmt struct {
	pgc        *postgresConn
	id         uint32
	paramCount int
}

func (stmt *postgreStmt) Close() error {
	/*
		if stmt.mc == nil || stmt.mc.closed.Load() {
			// driver.Stmt.Close can be called more than once, thus this function
			// has to be idempotent.
			// See also Issue #450 and golang/go#16019.
			//errLog.Print(ErrInvalidConn)
			return driver.ErrBadConn
		}

		err := stmt.mc.writeCommandPacketUint32(comStmtClose, stmt.id)
		stmt.mc = nil
		return err
	*/
	return nil
}
