package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"market.ir/dto"
	"market.ir/pkg/sdk"
)

func InitServer() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/q/:searching_keyword/nonce/:nonce_value", func(c *gin.Context) {
		//
		// Convert single keyword to multiple queries
		//
		var result []dto.TaggerArticleList
		key := c.Params.ByName("searching_keyword")
		q1 := sdk.Search(fmt.Sprintf("قیمت %s", key), "فروش")
		q2 := sdk.Search(fmt.Sprintf("بازار %s", key), "فروش")
		q3 := sdk.Search(fmt.Sprintf("سهم بازار %s", key), "فروش")

		result = append(result, q1, q2, q3)
		c.JSON(http.StatusOK, dto.ResponceDto{
			Request: dto.SearchRequest{
				Date:     time.Now().String(),
				Paramter: c.Params.ByName("searching_keyword"),
			},
			Data: result,
		})
	})

	return r
}
