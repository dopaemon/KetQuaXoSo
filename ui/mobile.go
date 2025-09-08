package ui

import (
	"fmt"
	"image/color"

	"KetQuaXoSo/internal/configs"
	_ "KetQuaXoSo/internal/rss"
	"KetQuaXoSo/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func BuildMobileUI(w fyne.Window) {
	bannerText := "xskt"
	banner := canvas.NewText(bannerText, color.White)
	banner.TextStyle = fyne.TextStyle{Bold: true}
	banner.TextSize = 36

	provinceSelect := widget.NewSelect(configs.Provinces, func(value string) {
		status.SetText("Đang tải dữ liệu...")
		go fetchResults(value)
	})
	provinceSelect.PlaceHolder = "Chọn loại vé số"

	dateSelect = widget.NewSelect([]string{}, func(value string) {
		showResults(value)
	})
	dateSelect.PlaceHolder = "Chọn ngày"

	resultsText = widget.NewMultiLineEntry()
	resultsText.SetMinRowsVisible(12)
	resultsText.Disable()

	input := widget.NewEntry()
	input.SetPlaceHolder("Nhập số cần kiểm tra")

	checkBtn := widget.NewButton("Kiểm tra", func() {
		youNum := input.Text
		if selectedDate == "" {
			dialog.ShowInformation("Thông báo", "Hãy chọn ngày trước", w)
			return
		}
		giai, num := utils.CheckWinningNumber(parsedResults, selectedDate, youNum)
		if giai != "" {
			dialog.ShowInformation("Kết quả",
				fmt.Sprintf("Số %s trúng giải %s: %s", youNum, giai, num), w)
		} else {
			dialog.ShowInformation("Kết quả", "Không trúng!", w)
		}
	})

	status = widget.NewLabel("")

	content := container.NewVBox(
		container.NewCenter(banner),
		provinceSelect,
		dateSelect,
		resultsText,
		input,
		checkBtn,
		status,
	)

	w.SetContent(container.NewScroll(content))
	w.Resize(fyne.NewSize(400, 700))
}
