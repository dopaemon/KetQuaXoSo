package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Version string
	Port    string
	Origins string
	ApiKey  string
}

var (
	AppConfig Config

	VERSION   = "v0.0.1"
	Provinces = []string{
		"Miền Bắc", "An Giang", "Bình Dương", "Bình Định", "Bạc Liêu", "Bình Phước",
		"Bến Tre", "Bình Thuận", "Cà Mau", "Cần Thơ", "Đắk Lắk", "Đồng Nai", "Đà Nẵng",
		"Đắk Nông", "Đồng Tháp", "Gia Lai", "Hồ Chí Minh", "Hậu Giang", "Kiên Giang",
		"Khánh Hòa", "Kon Tum", "Long An", "Lâm Đồng", "Ninh Thuận", "Phú Yên",
		"Quảng Bình", "Quảng Ngãi", "Quảng Nam", "Quảng Trị", "Sóc Trăng",
		"Tiền Giang", "Tây Ninh", "Thừa Thiên Huế", "Trà Vinh", "Vĩnh Long", "Vũng Tàu",
	}
	DateXS = []string{}
	Order  = []string{"ĐB", "1", "2", "3", "4", "5", "6", "7", "8"}

	Port    = "8080"
	Origins = "*"
)

func LoadConfig(envPath ...string) {
	var err error
	if len(envPath) > 0 {
		err = godotenv.Load(envPath[0])
		if err != nil {
			fmt.Printf("404: .env file not found in %s (Do not worry, default values loaded)\n", envPath[0])
		}
	} else {
		err = godotenv.Load()
		if err != nil {
			fmt.Println("404: .env file not found (Do not worry, default values loaded)")
		}
	}

	AppConfig = Config{
		Version: getEnv("APP_VERSION", VERSION),
		Port:    getEnv("APP_PORT", Port),
		Origins: getEnv("APP_ORIGINS", Origins),
		ApiKey:  getEnv("API_KEY", "YouAPIKey"),
	}

	Port = AppConfig.Port
	Origins = AppConfig.Origins
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
