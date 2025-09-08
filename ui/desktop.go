package ui

import (
	"image/color"

	"KetQuaXoSo/internal/configs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func BuildDesktopUI(w fyne.Window) {
	ui := &UIState{
		ResultsLabel: widget.NewLabel(""),
		Status:       widget.NewLabel(""),
	}

	banner := canvas.NewText("XSKT", color.White)
	banner.TextStyle = fyne.TextStyle{Bold: true}
	banner.TextSize = 50

	provinceSelect := widget.NewSelect(configs.Provinces, func(value string) {
		ui.Status.SetText("Đang tải dữ liệu...")
		ui.ResultsLabel.SetText("")
		go FetchResults(value, ui)
	})
	provinceSelect.PlaceHolder = "Chọn loại vé số"

	ui.DateSelect = widget.NewSelect([]string{}, func(value string) {
		ui.ResultsLabel.SetText("")
		ShowResults(value, ui)
	})
	ui.DateSelect.PlaceHolder = "Chọn ngày"

	input := widget.NewEntry()
	input.SetPlaceHolder("Nhập số cần kiểm tra")
	checkBtn := widget.NewButton("Kiểm tra", func() {
		CheckNumber(input.Text, ui, w)
	})

	left := container.NewVBox(
		container.NewCenter(banner),
		provinceSelect,
		ui.DateSelect,
		input,
		checkBtn,
		ui.Status,
	)

	right := container.NewMax(ui.ResultsLabel)

	content := container.NewHSplit(left, right)
	content.SetOffset(0.35)

	w.SetContent(content)
	w.Resize(fyne.NewSize(900, 600))
}
