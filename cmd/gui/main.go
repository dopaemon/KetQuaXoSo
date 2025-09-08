package main

import (
	"runtime"

	"KetQuaXoSo/ui"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.NewWithID("com.dopaemon.ketquaxoso")
	w := a.NewWindow("Kết Quả Xổ Số")

	if runtime.GOOS == "android" || runtime.GOOS == "ios" {
		ui.BuildMobileUI(w)
	} else {
		ui.BuildDesktopUI(w)
	}

	w.ShowAndRun()
}
