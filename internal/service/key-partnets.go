package service

import (
	"market.ir/internal/db"
	"market.ir/internal/dto"
	. "market.ir/pkg/text-minipulator"
)

func FindKeyPartnets(query string) (bool, []dto.Item) {
	art := &dto.Article{}
	result := db.Db.Find(&art)
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
