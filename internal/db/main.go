package db

import (
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Client() {
	var err error
	dsn := "tcp://localhost:8123?debug=true&database=content"
	// url := "tcp://" + config.Config["db_address"].(string) + ":9000?debug=true"
	Db, err = gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// if err := Db.Ping(); err != nil {
	// 	if exception, ok := err.(*clickhouse.Exception); ok {
	// 		fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
	// 	} else {
	// 		fmt.Println(err)
	// 	}
	// 	return
	// }
}
