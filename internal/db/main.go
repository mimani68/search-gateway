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
	// url := "tcp://" + config.Config["db_address"].(string) + ":9000?debug=true"
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
	// art := &dto.Article{}
	// result := Db.Find(&art)
	// // result := db.Where("language = ?", "fa").First(&dto.Article{})
	// // raws := db.First(&dto.Article, "language = ?", "fa")
	// // raws := db.Find(dto.Article{}, "body LIKE '%ترک%'")
	// // raws := db.Exec(`select * from content.article`)
	// if result.Error != nil {
	// 	panic("Error in retrive data")
	// }
	// fmt.Printf("[>>] %s", art)
	// fmt.Printf("[>>] %s", result)
	// if err := Db.Ping(); err != nil {
	// 	if exception, ok := err.(*clickhouse.Exception); ok {
	// 		fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
	// 	} else {
	// 		fmt.Println(err)
	// 	}
	// 	return
	// }
	// return Db
}
