package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ClickHouse/clickhouse-go"
)

var Db *sql.DB

func Client() {
	var err error
	Db, err = sql.Open("clickhouse", "tcp://db:9000?debug=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := Db.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
		return
	}
}
