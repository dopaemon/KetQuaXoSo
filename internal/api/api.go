package api

import (
	"net/http"
	"time"

	"KetQuaXoSo/internal/configs"
	"KetQuaXoSo/internal/rss"
	"KetQuaXoSo/utils"

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

type TicketRequest struct {
	Province string `json:"province"`
	Date     string `json:"date"`
	Number   string `json:"number"`
}

type TicketResponse struct {
	Province string `json:"province"`
	Date     string `json:"date"`
	Input    string `json:"input"`
	Prize    string `json:"prize"`
	Match    string `json:"match"`
	Error    string `json:"error,omitempty"`
}

func RunAPI() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{configs.Origins},
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

	r.POST("/api/check-ticket", func(c *gin.Context) {
		var req TicketRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}
		if req.Province == "" || req.Date == "" || req.Number == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "province, date, and number are required"})
			return
		}

		url := rss.Sources(req.Province)
		if url == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown province: " + req.Province})
			return
		}

		data, err := rss.Fetch(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, TicketResponse{
				Province: req.Province,
				Date:     req.Date,
				Input:    req.Number,
				Error:    "failed to fetch RSS: " + err.Error(),
			})
			return
		}

		results, err := rss.Parse(data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, TicketResponse{
				Province: req.Province,
				Date:     req.Date,
				Input:    req.Number,
				Error:    "failed to parse RSS: " + err.Error(),
			})
			return
		}

		prize, match := utils.CheckWinningNumber(results, req.Date, req.Number)

		c.JSON(http.StatusOK, TicketResponse{
			Province: req.Province,
			Date:     req.Date,
			Input:    req.Number,
			Prize:    prize,
			Match:    match,
		})
	})

	r.Run(":" + configs.Port)
}
