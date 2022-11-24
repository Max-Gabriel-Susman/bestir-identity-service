package database

type postgresConn struct {
}

func (pgc *postgresConn) Close() (err error) {
	/*
		// Makes Close idempotent
		if !mc.closed.Load() {
			err = mc.writeCommandPacket(comQuit)
		}

		mc.cleanup()
	*/
	return nil
}
