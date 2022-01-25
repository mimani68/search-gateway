package job

import (
	"time"

	"market.ir/pkg/scraping"
)

func trainTask() {
	keys := []string{
		"فروش",
		"قیمت",
		"سود",
		"سهم",
		"بازار",
		"بازده",
		"سرمایه گذاری",
		"استارتاپ",
		"نوآوری",
		"دانش بنیان",
		"محصول",
		"توسعه محصول",
		"نمونه اولیه",
		"مشتریان",
		"بازاریابی",
	}
	for _, key := range keys {
		go scraping.ScrapingHandler(key, "zoomit")
		go scraping.ScrapingHandler(key, "donya-e-eqtesad")
		go scraping.ScrapingHandler(key, "iribnews")
		go scraping.ScrapingHandler(key, "irna")
		go scraping.ScrapingHandler(key, "tabnak")
		go scraping.ScrapingHandler(key, "fars")
		time.Sleep(10 * time.Second)
	}
}
