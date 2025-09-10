//go:build headless

package main

import (
	"KetQuaXoSo/internal/configs"
	"KetQuaXoSo/internal/api"
)

func main() {
	configs.LoadConfig()
	api.RunAPI()
}
