package service

import (
	"fmt"

	"market.ir/internal/dto"
	textminipulator "market.ir/pkg/text-minipulator"
)

func FindKeyActivities(query string) (bool, []dto.Item) {
	fmt.Println(query)
	var text = `استاندار آذربایجان‌شرقی به جلسه خود با سفیر ایران در آنکارا اشاره کرد و گفت: هرگونه مانع در مسیر توسعه روابط با ترکیه از میان برداشته خواهد شد؛ البته موانع موجود بسیار کم و قابل حل هستند.
 
	وی از سرکنسول ترکیه در تبریز خواست، تسهیل تردد اتباع ایرانی، تجار و سرمایه‌گذاران از مرز بازرگان را پیگیری کند.
	 
	استاندار آذربایجان‌شرقی از آماده بودن زیرساخت‌ها در مراکز صنعتی چون شهرک سرمایه‌گذاری خارجی، منطقه ویژه اقتصادی سهلان، شهرک بعثت و نیز منطقه آزاد ارس برای سرمایه‌گذاری خبر داد و آمادگی مسوولان دولتی و بخش خصوصی آذربایجان‌شرقی برای تسهیل حضور سرمایه‌گذاران ترک را در استان اعلام کرد.`

	a := textminipulator.ShowText(text, query)

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
