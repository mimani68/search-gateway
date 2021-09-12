package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"market.ir/internal/dto"
	"market.ir/pkg/sdk"
)

func InitServer() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/q/:searching_keyword/nonce/:nonce_value", func(c *gin.Context) {
		//
		// Convert single keyword to multiple queries
		//
		var result []dto.TaggerArticleList
		key := c.Params.ByName("searching_keyword")
		q1 := sdk.Search(fmt.Sprintf("قیمت %s", key), "zoomit", "فروش")
		// q2 := sdk.Search(fmt.Sprintf("بازار %s", key), "iran", "سهم بازار")
		// q3 := sdk.Search(fmt.Sprintf("سهم بازار %s", key), "donya-e-eqtesad", "سهم بازار")

		result = append(result, q1)
		c.JSON(http.StatusOK, dto.ResponceDto{
			Request: dto.SearchRequest{
				Date:     time.Now().String(),
				Paramter: c.Params.ByName("searching_keyword"),
			},
			Data: result,
		})
	})

	r.GET("/dev/:param", func(c *gin.Context) {
		a := map[string]interface{}{
			"Date":     time.Now().String(),
			"Paramter": c.Params.ByName("param"),
			"Query":    c.Query("q"),
		}
		c.JSON(http.StatusOK, a)
	})

	return r
}
