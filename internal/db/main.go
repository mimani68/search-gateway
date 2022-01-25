package db

import (
	"os"

	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"market.ir/pkg/logHandler"
)

var Db *gorm.DB

func New() {
	var err error
	dsn := "tcp://" + os.Getenv("DB_ADDRESS") + ":" + os.Getenv("DB_PORT") + "?debug=true&database=" + os.Getenv("DB_NAME")
	Db, err = gorm.Open(clickhouse.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	} else {
		logHandler.Log("Database connected successfully.")
	}
}
