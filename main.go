//go:build !headless

package main

import (
	"os"
	"runtime"

	"github.com/dopaemon/KetQuaXoSo/internal/configs"
	"github.com/dopaemon/KetQuaXoSo/internal/api"
	"github.com/dopaemon/KetQuaXoSo/utils"
	"github.com/dopaemon/KetQuaXoSo/ui"

	"fyne.io/fyne/v2/app"
	_ "fyne.io/fyne/v2/theme"
)

func realMain() int {
	switch utils.GenFlags() {
		case "gui":
			a := app.NewWithID("com.dopaemon.ketquaxoso")
			w := a.NewWindow("Kết Quả Xổ Số")

			if runtime.GOOS == "android" || runtime.GOOS == "ios" {
				ui.BuildMobileUI(w)
			} else {
				ui.BuildDesktopUI(w)
			}
			w.ShowAndRun()
			return 0
		case "cli":
			ui.Tui()
			return 0
		case "api":
			configs.LoadConfig()
			api.RunAPI()
			return 0
		default:
			return 1
	}

	return 0
}

func main() {
	os.Exit(realMain())
}
