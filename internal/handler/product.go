package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"market.ir/internal/dto"
)

func ProductMetadataHandler(c *gin.Context) {
	category := []dto.ProductCategory{
		{
			Id:         "126510",
			Title:      "Home services",
			TitleFarsi: "لوازم خانگی",
		},
	}
	result := dto.ProductMetadataResponse{
		Category: category,
	}
	c.JSON(http.StatusOK, result)
}
