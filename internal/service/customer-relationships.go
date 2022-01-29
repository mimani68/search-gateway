package service

import (
	"market.ir/internal/db"
	"market.ir/internal/dto"
	textminipulator "market.ir/pkg/text-minipulator"
)

func FindCustomerRelationships(query string) (bool, []dto.Item) {
	var response []dto.Item
	list := []dto.Article{}
	// query = "suffix " + query + " suffix"
	dbQueryState := db.Db.Where("body LIKE ?", "% "+query+" %").Find(&list)
	if dbQueryState.Error != nil {
		panic("Error in retrive data")
	}
	for _, value := range list {
		a := textminipulator.ShowText(value.Body, query)
		if a["success"] != "true" {
			continue
		}
		response = append(response, dto.Item{
			Id:           value.Id,
			Title:        value.Title,
			PreText:      a["pre"].(string),
			PostText:     a["post"].(string),
			BoldText:     a["selectItem"].(string),
			CompleteText: a["complete"].(string),
			Link:         value.Link,
		})
	}
	return true, response
}
