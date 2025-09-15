//go:build headless

package main

import (
	"github.com/dopaemon/KetQuaXoSo/internal/configs"
	"github.com/dopaemon/KetQuaXoSo/internal/api"
)

func main() {
	configs.LoadConfig()
	api.RunAPI()
}
