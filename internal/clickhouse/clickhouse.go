package clickhouse

import (
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

func NewConn(addr []string) driver.Conn {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: addr,
		Auth: clickhouse.Auth{
			Database: "goweb",
			Username: "",
			Password: "",
		},
	})

	if err != nil {
		panic(err)
	}

	v, err := conn.ServerVersion()
	fmt.Println(v)

	return conn
}
