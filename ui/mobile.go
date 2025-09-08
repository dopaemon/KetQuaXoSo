package ui

import (
	"image/color"

	"KetQuaXoSo/internal/configs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func BuildMobileUI(w fyne.Window) {
	ui := &UIState{
		ResultsLabel: widget.NewLabel(""),
		Status:       widget.NewLabel(""),
	}

	banner := canvas.NewText("XSKT", color.White)
	banner.TextStyle = fyne.TextStyle{Bold: true}
	banner.TextSize = 36

	provinceSelect := widget.NewSelect(configs.Provinces, func(value string) {
		ui.Status.SetText("Đang tải dữ liệu...")
		go FetchResults(value, ui)
	})
	provinceSelect.PlaceHolder = "Chọn loại vé số"

	ui.DateSelect = widget.NewSelect([]string{}, func(value string) {
		ShowResults(value, ui)
	})
	ui.DateSelect.PlaceHolder = "Chọn ngày"

	input := widget.NewEntry()
	input.SetPlaceHolder("Nhập số cần kiểm tra")
	checkBtn := widget.NewButton("Kiểm tra", func() {
		CheckNumber(input.Text, ui, w)
	})

	content := container.NewVBox(
		container.NewCenter(banner),
		provinceSelect,
		ui.DateSelect,
		ui.ResultsLabel,
		input,
		checkBtn,
		ui.Status,
	)

	w.SetContent(container.NewScroll(content))
	w.Resize(fyne.NewSize(400, 700))
}
