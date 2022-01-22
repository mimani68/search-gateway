package server

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"market.ir/config"
	"market.ir/internal/db"
	"market.ir/internal/handler"
)

func InitServer() *gin.Engine {
	go config.GetConfig()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/ping", handler.Ping)
	r.GET("/q/:searching_keyword", handler.QueryHandler)

	go db.Client()

	return r
}
