package service

import "market.ir/internal/dto"

func ShowSingleQuery() (bool, dto.Item) {
	return false, dto.Item{
		Id:           "3753",
		Order:        1,
		PreText:      "متن در این قسمت به پیاان می شدر",
		PostText:     "در ادامه می توان به این اشاره کرد که ...",
		BoldText:     "میزان واردات سالیانه مناسب نیست",
		CompleteText: "sfsdf",
		Link:         "https://isna.ir/9849569",
	}
}
