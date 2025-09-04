package main

import (
	"fmt"
	"os"

	"XoSoToanQuoc/internal/rss"
	"XoSoToanQuoc/utils"
)

func main() {
	fmt.Println(utils.Banner())

	url := "https://xskt.com.vn/rss-feed/mien-nam-xsmn.rss"
	items, err := rss.Fetch(url)
	if err != nil {
		panic(err)
	}

	for _, item := range items {
		fmt.Println(item.Title)
		fmt.Println(item.Description)
		fmt.Println(item.Link)
		fmt.Println("-----------")
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
