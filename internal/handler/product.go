package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"market.ir/internal/dto"
)

func ProductMetadataHandler(c *gin.Context) {
	category := []dto.ProductCategory{
		{
			Id:    "126510",
			Title: "لوازم خانگی",
		},
	}
	metrcs := dto.Metrics{
		MarketCap:         "10000000",
		InterestForInvest: "0.2",
	}
	result := dto.ProductResponse{
		Category: category,
		Metrics:  metrcs,
	}
	c.JSON(http.StatusOK, result)
}
