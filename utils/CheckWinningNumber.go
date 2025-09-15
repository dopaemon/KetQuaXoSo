package utils

import (
	"strings"

	"github.com/dopaemon/KetQuaXoSo/internal/configs"
	"github.com/dopaemon/KetQuaXoSo/internal/rss"
)

func CheckWinningNumber(results []rss.Result, wdate, input string) (string, string) {
	for _, r := range results {
		if r.Date != wdate {
			continue
		}

		for _, giai := range configs.Order {
			if numbers, ok := r.Prizes[giai]; ok {
				for _, num := range numbers {
					if strings.HasSuffix(input, num) {
						return giai, num
					}
				}
			}
		}
		break
	}

	return "", ""
}
