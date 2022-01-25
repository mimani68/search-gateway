package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"market.ir/internal/dto"
	"market.ir/internal/service"
	"market.ir/pkg/logHandler"
)

func QueryHandler(c *gin.Context) {
	query := c.Params.ByName("searching_keyword")

	//
	// FIXME: Store search request in database
	//

	var result dto.ResponseDto

	findKeyPartnersSuccess, keyPartnersItems := service.FindKeyPartnets(query)
	if !findKeyPartnersSuccess {
		logHandler.Log("Key Partners list for query: \"" + query + "\" was not exists.")
	}
	result.KeyPartners = keyPartnersItems

	findKeyActivitiesSuccess, keyActivitiesItems := service.FindKeyActivities(query)
	if !findKeyActivitiesSuccess {
		logHandler.Log("Key Activities list for query: \"" + query + "\" was not exists.")
	}
	result.KeyActivities = keyActivitiesItems

	findValuePropositionsSuccess, ValuePropositionsItems := service.FindValuePropositions(query)
	if !findValuePropositionsSuccess {
		logHandler.Log("Value Propositions list for query: \"" + query + "\" was not exists.")
	}
	result.ValuePropositions = ValuePropositionsItems

	findCustomerRelationshipsSuccess, CustomerRelationshipsItems := service.FindCustomerRelationships(query)
	if !findCustomerRelationshipsSuccess {
		logHandler.Log("Customer Relationships list for query: \"" + query + "\" was not exists.")
	}
	result.CustomerRelationships = CustomerRelationshipsItems

	findChannelsSuccess, ChannelsItems := service.FindChannels(query)
	if !findChannelsSuccess {
		logHandler.Log("Channel list for query: \"" + query + "\" was not exists.")
	}
	result.Channels = ChannelsItems

	findCustomerSegmentsSuccess, CustomerSegmentsItems := service.FindCustomerSegments(query)
	if !findCustomerSegmentsSuccess {
		logHandler.Log("Customer Segments list for query: \"" + query + "\" was not exists.")
	}
	result.CustomerSegments = CustomerSegmentsItems

	findCostStructureSuccess, CostStructureItems := service.FindCostStructures(query)
	if !findCostStructureSuccess {
		logHandler.Log("Cost Structure list for query: \"" + query + "\" was not exists.")
	}
	result.CostStructure = CostStructureItems

	findRevenueStreamsSuccess, RevenueStreamsItems := service.FindRevenueStreams(query)
	if !findRevenueStreamsSuccess {
		logHandler.Log("Revenue Streams list for query: \"" + query + "\" was not exists.")
	}
	result.RevenueStreams = RevenueStreamsItems

	result.Meta.Paramter = query

	c.JSON(http.StatusOK, result)

}
