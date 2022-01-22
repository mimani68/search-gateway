package config

import "os"

var Config map[string]interface{}

func GetConfig(configKey string) {
	Config = map[string]interface{}{
		"SCRAPER":      "http://scraper:3000",
		"resultNumber": "50",
		"db_address":   os.Getenv("DB_ADDRESS"),
		"db_port":      os.Getenv("DB_PORT"),
		"db_username":  os.Getenv("DB_USERNAME"),
		"db_password":  os.Getenv("DB_PASSWORD"),
	}
}
