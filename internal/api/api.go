package api

import (
	"time"
	"net/http"

	"KetQuaXoSo/internal/configs"
	"KetQuaXoSo/internal/rss"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type CheckRequest struct {
	Province string `json:"province"`
}

type CheckResponse struct {
	Province string      `json:"province"`
	Results  interface{} `json:"results"`
	Error    string      `json:"error,omitempty"`
}

func RunAPI() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
		MaxAge:       12 * time.Hour,
	}))

	r.GET("/api/province", func(c *gin.Context) {
		c.JSON(http.StatusOK, configs.Provinces)
	})

	r.POST("/api/check", func(c *gin.Context) {
		var req CheckRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}
		if req.Province == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "province is required"})
			return
		}

		url := rss.Sources(req.Province)
		if url == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown province: " + req.Province})
			return
		}

		data, err := rss.Fetch(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, CheckResponse{
				Province: req.Province,
				Error:    "failed to fetch RSS: " + err.Error(),
			})
			return
		}

		results, err := rss.Parse(data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, CheckResponse{
				Province: req.Province,
				Error:    "failed to parse RSS: " + err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, CheckResponse{
			Province: req.Province,
			Results:  results,
		})
	})

	r.Run(":" + configs.Port)
}
