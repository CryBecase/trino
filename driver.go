package trino

import "database/sql/driver"

// sqldriver implements driver.Driver
type sqldriver struct{}

func (d *sqldriver) Open(name string) (driver.Conn, error) {
	return newConn(name)
}

var _ driver.Driver = &sqldriver{}
