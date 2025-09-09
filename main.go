// @title KetQuaXoSo API
// @version 1.0
// @description API kiểm tra kết quả xổ số
// @contact.name dopaemon
// @contact.email polarisdp@gmail.com
// @host localhost:8080
// @BasePath /

package main

import (
	"os"
	"runtime"

	"KetQuaXoSo/internal/api"
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
		case "api":
			api.RunAPI()
			break
		default:
			return 1
	}

	return 0
}

func main() {
	os.Exit(realMain())
}
