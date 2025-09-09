package main

import (
	"os"
	"runtime"

	"KetQuaXoSo/utils"
	"KetQuaXoSo/ui"

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
			break
		case "cli":
			ui.Tui()
			break
		default:
			return 1
	}

	return 0
}

func main() {
	os.Exit(realMain())
}
