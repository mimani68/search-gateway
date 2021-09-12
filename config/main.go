package config

func GetConfig(configKey string) string {
	configPool := map[string]interface{}{
		"SCRAPER":      "http://scraper:3000",
		"resultNumber": "50",
	}
	return configPool[configKey].(string)
}
