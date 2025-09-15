package ui

import (
	"fmt"
	"os"

	"github.com/dopaemon/KetQuaXoSo/internal/configs"
	"github.com/dopaemon/KetQuaXoSo/internal/rss"
	"github.com/dopaemon/KetQuaXoSo/utils"

	"github.com/charmbracelet/huh"
)

func Tui() {
	fmt.Println(utils.Banner())

	province := ""
	wdate := ""

	provinceOptions := make([]huh.Option[string], len(configs.Provinces))
	for i, code := range configs.Provinces {
		provinceOptions[i] = huh.NewOption(code, code)
	}

	provinceSelect := huh.NewSelect[string]().
		Title("Chọn loại vé số:").
		Options(provinceOptions...).
		Value(&province)

	huh.NewForm(huh.NewGroup(provinceSelect)).Run()

	url, _ := rss.Sources(province)
	data, err := rss.Fetch(url)
	if err != nil {
		fmt.Println("Lỗi fetch RSS:", err)
		os.Exit(1)
	}

	results, err := rss.Parse(data)
	if err != nil {
		fmt.Println("Lỗi parse RSS:", err)
		os.Exit(1)
	}

	dateOptions := make([]huh.Option[string], len(results))
	for i, r := range results {
		dateOptions[i] = huh.NewOption(r.Date, r.Date)
	}

	dateSelect := huh.NewSelect[string]().
		Title("Chọn ngày:").
		Options(dateOptions...).
		Value(&wdate)

	huh.NewForm(huh.NewGroup(dateSelect)).Run()

	found := false
	for _, r := range results {
		if r.Date == wdate {
			fmt.Println("=== ", r.Title, " ===")
			for _, giai := range configs.Order {
				if so, ok := r.Prizes[giai]; ok {
					fmt.Println("Giải", giai+":", so)
				}
			}
			found = true
			break
		}
	}

	if !found {
		fmt.Println("!!! Không có kết quả cho ngày này !!!")
		os.Exit(0)
	}

	youNum := ""
	input := huh.NewInput().
		Title("Nhập số để kiểm tra").
		Prompt(": ").
		Value(&youNum)
	huh.NewForm(huh.NewGroup(input)).Run()

	giai, num := utils.CheckWinningNumber(results, wdate, youNum)
	if giai != "" {
		fmt.Printf("\n\nSố %s của bạn là số trúng! Giải %s: %s\n", youNum, giai, num)
	} else {
		fmt.Println("\n\nKhông trúng!")
	}
}
