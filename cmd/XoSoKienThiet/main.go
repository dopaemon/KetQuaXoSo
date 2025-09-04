package main

import (
	"fmt"
	"os"

	"XoSoToanQuoc/utils"
)

func main() {
	fmt.Println(utils.Banner())

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
