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
}

var (
	AppConfig Config

	VERSION	string = "v0.0.1"
	Provinces []string = []string{"Miền Bắc", "An Giang", "Bình Dương", "Bình Định", "Bạc Liêu", "Bình Phước", "Bến Tre", "Bình Thuận", "Cà Mau", "Cần Thơ", "Đắk Lắk", "Đồng Nai", "Đà Nẵng", "Đắk Nông", "Đồng Tháp", "Gia Lai", "Hồ Chí Minh", "Hậu Giang", "Kiên Giang", "Khánh Hòa", "Kon Tum", "Long An", "Lâm Đồng", "Ninh Thuận", "Phú Yên", "Quảng Bình", "Quảng Ngãi", "Quảng Nam", "Quảng Trị", "Sóc Trăng", "Tiền Giang", "Tây Ninh", "Thừa Thiên Huế", "Trà Vinh", "Vĩnh Long", "Vũng Tàu"}
	DateXS []string = []string{}
	Order []string = []string{"ĐB", "1", "2", "3", "4", "5", "6", "7", "8"}
	Port string = "8080"
	Origins string = "*"
)

func LoadConfig(envPath ...string) {
	if len(envPath) > 0 {
		if err := godotenv.Load(envPath[0]); err != nil {
			fmt.Printf("404: .env file not found in %s (Do not worrie, default values loaded)\n", envPath[0])
		}
	} else {
		if err := godotenv.Load(); err != nil {
			fmt.Println("404: .env file not found (Do not worrie, default values loaded)")
		}
	}

	AppConfig = Config{
		Version: getEnv("APP_VERSION", VERSION),
		Port:    getEnv("APP_PORT", Port),
		Origins: getEnv("APP_ORIGINS", Origins),
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
