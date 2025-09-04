package main

import (
	"fmt"
	"os"

	"XoSoToanQuoc/internal/configs"
	"XoSoToanQuoc/internal/rss"
	"XoSoToanQuoc/utils"

	"github.com/charmbracelet/huh"
)

func main() {
	fmt.Println(utils.Banner())

	var province string

	input := huh.NewInput().
		Title("Bạn muốn xem kết quả sổ số nào ?").
		Prompt(": ").
		Suggestions(configs.Provinces).
		Value(&province)

	huh.NewForm(huh.NewGroup(input)).Run()

	url := rss.Sources(province)
	data, _ := rss.Fetch(url)
	results, _ := rss.Parse(data)

	for _, r := range results {
		fmt.Println("=== ", r.Title, " ===")
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
