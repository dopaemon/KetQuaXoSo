package main

import (
	"os"
	"runtime"

	"KetQuaXoSo/utils"
	"KetQuaXoSo/ui"

	"fyne.io/fyne/v2/app"
	_ "fyne.io/fyne/v2/theme"
)

func main() {
	a := app.NewWithID("com.dopaemon.ketquaxoso")
	w := a.NewWindow("Kết Quả Xổ Số")

	if runtime.GOOS == "android" || runtime.GOOS == "ios" {
		ui.BuildMobileUI(w)
	} else {
		ui.BuildDesktopUI(w)
	}

	switch utils.GenFlags() {
		case "gui":
			w.ShowAndRun()
			break
		case "cli":
			ui.Tui()
			break
		default:
			os.Exit(1)
	}
}

