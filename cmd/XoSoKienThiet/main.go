package main

import (
	"fmt"
	"os"

	"XoSoToanQuoc/internal/rss"
	"XoSoToanQuoc/utils"
)

func main() {
	fmt.Println(utils.Banner())

	url := rss.Sources("an-giang")
	data, _ := rss.Fetch(url)
	results, _ := rss.Parse(data)

	for _, r := range results {
		fmt.Println("=== ", r.Province, r.Date, " ===")
		for giai, so := range r.Prizes {
			fmt.Println("Giải", giai+":", so)
		}
		fmt.Println()
	}

	switch utils.GenFlags() {
		case "gui":
			fmt.Println("Run GUI")
			break
		case "cli":
			fmt.Println("Run CLI")
			break
		default:
			os.Exit(1)
	}
}
