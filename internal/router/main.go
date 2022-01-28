package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"market.ir/internal/handler"
)

func RegisterRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/ping", handler.Ping)

	r.GET("/product/metadata", handler.ProductMetadataHandler)

	r.POST("/search", handler.QueryHandler)

	r.POST("/feedback", handler.FeedBackHandler)

	return r
}
