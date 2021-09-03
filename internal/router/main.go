package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"market.ir/dto"
	"market.ir/pkg/sdk"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/q/:searching_keyword/nonce/:nonce_value", func(c *gin.Context) {
		result := sdk.Search(c.Params.ByName("searching_keyword"))
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
