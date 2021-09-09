package server

import (
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
		// Search orginal
		//
		result := sdk.Search(c.Params.ByName("searching_keyword"))
		//
		// TODO: custom search
		//
		// marketCanvasA1Question := sdk.Search(c.Params.ByName("searching_keyword"))
		// marketCanvasA2Question := sdk.Search(c.Params.ByName("searching_keyword"))
		// marketCanvasA3Question := sdk.Search(c.Params.ByName("searching_keyword"))
		// marketCanvasB1Question := sdk.Search(c.Params.ByName("searching_keyword"))
		// marketCanvasB2Question := sdk.Search(c.Params.ByName("searching_keyword"))
		// marketCanvasB4Question := sdk.Search(c.Params.ByName("searching_keyword"))
		// marketCanvasC1Question := sdk.Search(c.Params.ByName("searching_keyword"))
		// marketCanvasC2Question := sdk.Search(c.Params.ByName("searching_keyword"))

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
