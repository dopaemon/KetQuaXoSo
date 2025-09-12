package api

import (
	"net/http"
	"time"

	"KetQuaXoSo/internal/configs"
	"KetQuaXoSo/internal/rss"
	"KetQuaXoSo/utils"

	_ "KetQuaXoSo/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title KetQuaXoSo API
// @version 1.0
// @description API tra cứu kết quả xổ số kiến thiết Việt Nam.
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-Key
// @description Nhập API key để sử dụng API. Giá trị mặc định cho test: **YouAPIKey**

type CheckRequest struct {
	Province string `json:"province" example:"Lâm Đồng"`
}

type CheckResponse struct {
	Province string      `json:"province"`
	Results  interface{} `json:"results"`
	Error    string      `json:"error,omitempty"`
}

type TicketRequest struct {
	Province string `json:"province" example:"Lâm Đồng"`
	Date     string `json:"date" example:"07/09"`
	Number   string `json:"number" example:"123456"`
}

type TicketResponse struct {
	Province string `json:"province"`
	Date     string `json:"date"`
	Input    string `json:"input"`
	Prize    string `json:"prize"`
	Match    string `json:"match"`
	Error    string `json:"error,omitempty"`
}

func APIKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")
		expectedKey := configs.AppConfig.ApiKey

		if expectedKey != "" && apiKey != expectedKey {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing API key"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// @Summary	Lấy danh sách tỉnh
// @Description	Trả về danh sách các tỉnh có hỗ trợ xổ số
// @Tags	Province
// @Security	ApiKeyAuth
// @Produce	json
// @Success	200 {array} string
// @Failure	401 {object} map[string]string
// @Router	/api/province [get]
func GetProvinces(c *gin.Context) {
	c.JSON(http.StatusOK, configs.Provinces)
}

// @Summary	Lấy kết quả xổ số
// @Description	Lấy toàn bộ kết quả xổ số của một tỉnh
// @Tags	Lottery
// @Security	ApiKeyAuth
// @Accept	json
// @Produce	json
// @Param	request body CheckRequest true "Thông tin tỉnh"
// @Success	200 {object} CheckResponse
// @Failure	400 {object} CheckResponse
// @Failure	401 {object} map[string]string
// @Failure	500 {object} CheckResponse
// @Router	/api/check [post]
func CheckLottery(c *gin.Context) {
	var req CheckRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, CheckResponse{Error: "Invalid JSON"})
		return
	}
	if req.Province == "" {
		c.JSON(http.StatusBadRequest, CheckResponse{Error: "province is required"})
		return
	}

	url, _ := rss.Sources(req.Province)
	if url == "" {
		c.JSON(http.StatusBadRequest, CheckResponse{Error: "Unknown province: " + req.Province})
		return
	}

	data, err := rss.Fetch(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CheckResponse{Province: req.Province, Error: err.Error()})
		return
	}

	results, err := rss.Parse(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CheckResponse{Province: req.Province, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, CheckResponse{
		Province: req.Province,
		Results:  results,
	})
}

// @Summary	Kiểm tra vé số
// @Description	Kiểm tra xem vé số có trúng thưởng hay không
// @Tags	Lottery
// @Security	ApiKeyAuth
// @Accept	json
// @Produce	json
// @Param	request body TicketRequest true "Thông tin vé số"
// @Success	200 {object} TicketResponse
// @Failure	400 {object} TicketResponse
// @Failure	401 {object} map[string]string
// @Failure	500 {object} TicketResponse
// @Router	/api/check-ticket [post]
func CheckTicket(c *gin.Context) {
	var req TicketRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, TicketResponse{Error: "Invalid JSON"})
		return
	}
	if req.Province == "" || req.Date == "" || req.Number == "" {
		c.JSON(http.StatusBadRequest, TicketResponse{Error: "province, date, and number are required"})
		return
	}

	url, _ := rss.Sources(req.Province)
	if url == "" {
		c.JSON(http.StatusBadRequest, TicketResponse{Error: "Unknown province: " + req.Province})
		return
	}

	data, err := rss.Fetch(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, TicketResponse{
			Province: req.Province,
			Date:     req.Date,
			Input:    req.Number,
			Error:    err.Error(),
		})
		return
	}

	results, err := rss.Parse(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, TicketResponse{
			Province: req.Province,
			Date:     req.Date,
			Input:    req.Number,
			Error:    err.Error(),
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
}

func RunAPI() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{configs.Origins},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "X-API-Key"},
		MaxAge:       12 * time.Hour,
	}))

	api := r.Group("/api")
	api.Use(APIKeyAuth())
	{
		api.GET("/province", GetProvinces)
		api.POST("/check", CheckLottery)
		api.POST("/check-ticket", CheckTicket)
	}

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.PersistAuthorization(true))) /* */

	r.Run(":" + configs.Port)
}
