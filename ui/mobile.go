package ui

import (
	"image/color"

	"github.com/dopaemon/KetQuaXoSo/internal/configs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func BuildMobileUI(w fyne.Window) {
	ui := &UIState{
		ResultsLabel: widget.NewLabel(""),
		Status:       widget.NewLabel(""),
		LinkRSS:      widget.NewHyperlink("", nil),
	}

	ui.ResultsLabel.Wrapping = fyne.TextWrapWord

	banner := canvas.NewText("KQXS", color.White)
	banner.TextStyle = fyne.TextStyle{Bold: true}
	banner.TextSize = 50

	des := canvas.NewText("Chương trình xem kết quả xổ số", color.White)
	des.TextStyle = fyne.TextStyle{Bold: true}
	des.TextSize = 20

	provinceSelect := widget.NewSelect(configs.Provinces, func(value string) {
		ui.Status.SetText("Đang tải dữ liệu...")
		ui.ResultsLabel.Hide()
		go FetchResults(value, ui)
	})
	provinceSelect.PlaceHolder = "Chọn loại vé số"

	ui.DateSelect = widget.NewSelect([]string{}, func(value string) {
		ui.ResultsLabel.Show()
		ShowResults(value, ui)
	})
	ui.DateSelect.PlaceHolder = "Chọn ngày"

	input := widget.NewEntry()
	input.SetPlaceHolder("Nhập số cần kiểm tra")
	checkBtn := widget.NewButton("Kiểm tra", func() {
		if !IsSixDigitNumber(input.Text) {
			ui.Status.SetText("Vé số phải có 5-6 ký tự số !!!")
			return
		}

		ui.Status.SetText("")
		CheckNumber(input.Text, ui, w)
	})

	hSep := canvas.NewRectangle(color.NRGBA{R: 200, G: 200, B: 200, A: 255})
	hSep.SetMinSize(fyne.NewSize(0, 1))

	row := container.NewBorder(nil, nil, nil, checkBtn, input)

	row1 := container.NewGridWithColumns(2,
		container.NewMax(ui.Status),
		container.NewHBox(
			layout.NewSpacer(),
			ui.LinkRSS,
		),
	)

	content := container.NewVBox(
		container.NewCenter(banner),
		container.NewCenter(des),
		provinceSelect,
		ui.DateSelect,
		row,
		ui.ResultsLabel,
		hSep,
		row1,
	)

	w.SetContent(container.NewScroll(content))
	w.Resize(fyne.NewSize(400, 700))
}
