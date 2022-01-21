package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"market.ir/internal/dto"
	"market.ir/internal/service"
)

func QueryHandler(c *gin.Context) {
	key := c.Params.ByName("searching_keyword")

	//
	// Call for automatic traning
	//
	// q1 := sdk.Search(fmt.Sprintf("قیمت %s", key), "zoomit", "فروش")

	var result dto.ResponseDto
	success, keyPartnersItem := service.ShowSingleQuery()
	if !success {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Error": "true",
		})
		return
	}
	result.KeyPartners = append(result.KeyPartners, keyPartnersItem)
	result.Meta.Paramter = key

	c.JSON(http.StatusOK, result)

	// c.JSON(http.StatusOK, dto.ResponseDto{
	// 	Meta: dto.SearchRequest{
	// 		Paramter: key,
	// 	},
	// 	KeyPartners: []dto.Item{
	// 		result,
	// 	},
	// 	KeyActivities:         []dto.Item{},
	// 	KeyResources:          []dto.Item{},
	// 	ValuePropositions:     []dto.Item{},
	// 	CustomerRelationships: []dto.Item{},
	// 	Channels:              []dto.Item{},
	// 	CustomerSegments:      []dto.Item{},
	// 	CostStructure:         []dto.Item{},
	// 	RevenueStreams:        []dto.Item{},
	// })
}
