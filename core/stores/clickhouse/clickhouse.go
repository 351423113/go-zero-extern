package clickhouse

import (
	// imports the driver.
	"github.com/351423113/go-zero-extern/core/stores/sqlx"
	_ "github.com/ClickHouse/clickhouse-go"
)

const clickHouseDriverName = "clickhouse"

// New returns a clickhouse connection.
func New(datasource string, opts ...sqlx.SqlOption) sqlx.SqlConn {
	return sqlx.NewSqlConn(clickHouseDriverName, datasource, opts...)
}
