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
	var wdate string

	options := make([]huh.Option[string], len(configs.Provinces))
	for i, code := range configs.Provinces {
		options[i] = huh.NewOption(code, code)
	}

	selected := huh.NewSelect[string]().
	Title("Chọn loại vé số bạn muốn: ").
	Options(options...).
	Value(&province)

	huh.NewForm(huh.NewGroup(selected)).Run()

	url := rss.Sources(province)
	data, _ := rss.Fetch(url)
	results, _ := rss.Parse(data)

	for _, r := range results {
		configs.DateXS = append(configs.DateXS, r.Date)
	}

	for _, r := range results {
		input := huh.NewInput().
			Title("Bạn muốn xem kết quả sổ số ngày nào ?").
			Prompt(": ").
			Suggestions(configs.DateXS).
			Value(&wdate)
		huh.NewForm(huh.NewGroup(input)).Run()

		if wdate == r.Date {
			fmt.Println("=== ", r.Title, " ===")
			for giai, so := range r.Prizes {
				fmt.Println("Giải", giai+":", so)
			}
			break
		} else {
			fmt.Println("!!! Không có kết quả cho ngày này !!!")
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
