package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FeedBackHandler(c *gin.Context) {
	// key := c.Params.ByName("searching_keyword")

	c.JSON(http.StatusOK, map[string]interface{}{
		"success": "true",
	})

}
