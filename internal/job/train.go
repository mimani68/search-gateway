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
		"سرمایه گذاری خطر پذیر",
		"استارتاپ",
		"نوآوری",
		"دانش بنیان",
		"محصول",
		"رقابت",
		"کیفیت",
		"توسعه محصول",
		"نمونه اولیه",
		"مشتریان",
		"بازاریابی",
		"مجوز",
		"اتحادیه",
		"اصناف",
		"لایحه",
		"قانون",
		"نوآورانه",
		"معاونت علمی",
		"مراکز رشد فناوری",
		"پارک علمی",
		"تجاری سازی",
		"نرم افزار",
		"برنامه",
		"اپلیکیشن",
		"توسعه",
		"توسعه دهندگان",
		"هوش مصنوعی",
		"زیر ساخت",
		"فناوری",
		"تکنولوژی",
		"رمزارز",
		"ارز",
	}
	for _, key := range keys {
		go scraping.ScrapingHandler(key, "zoomit")
		go scraping.ScrapingHandler(key, "donya-e-eqtesad")
		go scraping.ScrapingHandler(key, "iribnews")
		go scraping.ScrapingHandler(key, "irna")
		go scraping.ScrapingHandler(key, "tabnak")
		go scraping.ScrapingHandler(key, "fars")
		time.Sleep(30 * time.Second)
	}
}
