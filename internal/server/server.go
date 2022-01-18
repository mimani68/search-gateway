package server

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"market.ir/internal/dto"
)

func InitServer() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/q/:searching_keyword", func(c *gin.Context) {
		//
		// Convert single keyword to multiple queries
		//
		// var result []dto.TaggerArticleList
		// key := c.Params.ByName("searching_keyword")
		// q1 := sdk.Search(fmt.Sprintf("قیمت %s", key), "zoomit", "فروش")
		// // q2 := sdk.Search(fmt.Sprintf("بازار %s", key), "iran", "سهم بازار")
		// // q3 := sdk.Search(fmt.Sprintf("سهم بازار %s", key), "donya-e-eqtesad", "سهم بازار")
		// result = append(result, q1)

		c.JSON(http.StatusOK, dto.ResponseDto{
			Meta: dto.SearchRequest{},
			KeyPartners: []dto.Item{
				{
					Id:           "3754",
					Order:        1,
					PreText:      "متن در این قسمت به پیاان می شدر",
					PostText:     "در ادامه می توان به این اشاره کرد که ...",
					BoldText:     "میزان واردات سالیانه مناسب نیست",
					CompleteText: "sfsdf",
					Link:         "https://isna.ir/9849569",
				},
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
