package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"market.ir/internal/dto"
	"market.ir/internal/service"
	"market.ir/pkg/logHandler"
)

func QueryHandler(c *gin.Context) {
	body := dto.SearchMeta{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//
	// FIXME: Store search request in database
	//

	var result dto.ResponseDto

	findKeyPartnersSuccess, keyPartnersItems := service.FindKeyPartnets(body.Query)
	if !findKeyPartnersSuccess {
		logHandler.Log("Key Partners list for query: \"" + body.Query + "\" was not exists.")
	}
	result.KeyPartners = keyPartnersItems

	findKeyActivitiesSuccess, keyActivitiesItems := service.FindKeyActivities(body.Query)
	if !findKeyActivitiesSuccess {
		logHandler.Log("Key Activities list for query: \"" + body.Query + "\" was not exists.")
	}
	result.KeyActivities = keyActivitiesItems

	findValuePropositionsSuccess, ValuePropositionsItems := service.FindValuePropositions(body.Query)
	if !findValuePropositionsSuccess {
		logHandler.Log("Value Propositions list for query: \"" + body.Query + "\" was not exists.")
	}
	result.ValuePropositions = ValuePropositionsItems

	findCustomerRelationshipsSuccess, CustomerRelationshipsItems := service.FindCustomerRelationships(body.Query)
	if !findCustomerRelationshipsSuccess {
		logHandler.Log("Customer Relationships list for query: \"" + body.Query + "\" was not exists.")
	}
	result.CustomerRelationships = CustomerRelationshipsItems

	findChannelsSuccess, ChannelsItems := service.FindChannels(body.Query)
	if !findChannelsSuccess {
		logHandler.Log("Channel list for query: \"" + body.Query + "\" was not exists.")
	}
	result.Channels = ChannelsItems

	findCustomerSegmentsSuccess, CustomerSegmentsItems := service.FindCustomerSegments(body.Query)
	if !findCustomerSegmentsSuccess {
		logHandler.Log("Customer Segments list for query: \"" + body.Query + "\" was not exists.")
	}
	result.CustomerSegments = CustomerSegmentsItems

	findCostStructureSuccess, CostStructureItems := service.FindCostStructures(body.Query)
	if !findCostStructureSuccess {
		logHandler.Log("Cost Structure list for query: \"" + body.Query + "\" was not exists.")
	}
	result.CostStructure = CostStructureItems

	findRevenueStreamsSuccess, RevenueStreamsItems := service.FindRevenueStreams(body.Query)
	if !findRevenueStreamsSuccess {
		logHandler.Log("Revenue Streams list for query: \"" + body.Query + "\" was not exists.")
	}
	result.RevenueStreams = RevenueStreamsItems

	result.Meta.Paramter = body.Query

	c.JSON(http.StatusOK, result)

}
