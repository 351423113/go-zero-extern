package postgres

import (
	// imports the driver.
	"github.com/351423113/go-zero-extern/core/stores/sqlx"
	_ "github.com/lib/pq"
)

const postgresDriverName = "postgres"

// New returns a postgres connection.
func New(datasource string, opts ...sqlx.SqlOption) sqlx.SqlConn {
	return sqlx.NewSqlConn(postgresDriverName, datasource, opts...)
}
