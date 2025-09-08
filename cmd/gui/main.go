package main

import (
	"runtime"

	"KetQuaXoSo/ui"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func main() {
	a := app.NewWithID("com.dopaemon.ketquaxoso")
	w := a.NewWindow("Kết Quả Xổ Số")
	a.Settings().SetTheme(theme.DarkTheme())

	if runtime.GOOS == "android" || runtime.GOOS == "ios" {
		ui.BuildMobileUI(w)
	} else {
		ui.BuildDesktopUI(w)
	}

	w.ShowAndRun()
}
