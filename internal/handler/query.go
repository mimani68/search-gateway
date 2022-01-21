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

	err, result := service.ShowSingleQuery()
	if err {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Error": "true",
		})
	}

	c.JSON(http.StatusOK, dto.ResponseDto{
		Meta: dto.SearchRequest{
			Paramter: key,
		},
		KeyPartners: []dto.Item{
			result,
		},
		KeyActivities:         []dto.Item{},
		KeyResources:          []dto.Item{},
		ValuePropositions:     []dto.Item{},
		CustomerRelationships: []dto.Item{},
		Channels:              []dto.Item{},
		CustomerSegments:      []dto.Item{},
		CostStructure:         []dto.Item{},
		RevenueStreams:        []dto.Item{},
	})
}
