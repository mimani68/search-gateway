package service

import (
	"market.ir/internal/db"
	"market.ir/internal/dto"
	. "market.ir/pkg/text-minipulator"
)

func FindChannels(query string) (bool, []dto.Item) {
	// var text = `استاندار آذربایجان‌شرقی به جلسه خود با سفیر ایران در آنکارا اشاره کرد و گفت: هرگونه مانع در مسیر توسعه روابط با ترکیه از میان برداشته خواهد شد؛ البته موانع موجود بسیار کم و قابل حل هستند.
	// وی از سرکنسول ترکیه در تبریز خواست، تسهیل تردد اتباع ایرانی، تجار و سرمایه‌گذاران از مرز بازرگان را پیگیری کند.
	// استاندار آذربایجان‌شرقی از آماده بودن زیرساخت‌ها در مراکز صنعتی چون شهرک سرمایه‌گذاری خارجی، منطقه ویژه اقتصادی سهلان، شهرک بعثت و نیز منطقه آزاد ارس برای سرمایه‌گذاری خبر داد و آمادگی مسوولان دولتی و بخش خصوصی آذربایجان‌شرقی برای تسهیل حضور سرمایه‌گذاران ترک را در استان اعلام کرد.`

	//
	// sql query
	//
	// raws := db.Db.Find(&dto.Article{}, "body LIKE '%?%'", query)
	// fmt.Println(raws)
	art := &dto.Article{}
	result := db.Db.Find(&art)
	// result := db.Where("language = ?", "fa").First(&dto.Article{})
	// raws := db.First(&dto.Article, "language = ?", "fa")
	// raws := db.Find(dto.Article{}, "body LIKE '%ترک%'")
	// raws := db.Exec(`select * from content.article`)
	if result.Error != nil {
		panic("Error in retrive data")
	}

	a := ShowText(art.Body, query)

	if a["success"] != "true" {
		return false, []dto.Item{}
	}

	return true, []dto.Item{
		{
			Id:           "3753",
			Order:        1,
			PreText:      a["pre"].(string),
			PostText:     a["post"].(string),
			BoldText:     a["selectItem"].(string),
			CompleteText: a["complete"].(string),
			Link:         "https://isna.ir/9849569",
		},
	}
}
