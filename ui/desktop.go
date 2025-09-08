package ui

import (
	"fmt"
	"image/color"

	"KetQuaXoSo/internal/configs"
	"KetQuaXoSo/internal/rss"
	"KetQuaXoSo/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var (
	dateSelectDesktop    *widget.Select
	resultsLabelDesktop  *widget.Label
	statusDesktop        *widget.Label
	parsedResultsDesktop []rss.Result
	selectedDateDesktop  string
)

func BuildDesktopUI(w fyne.Window) {
	banner := canvas.NewText("xskt", color.White)
	banner.TextStyle = fyne.TextStyle{Bold: true}
	banner.TextSize = 50

	provinceSelect := widget.NewSelect(configs.Provinces, func(value string) {
		statusDesktop.SetText("Đang tải dữ liệu...")
		go fetchResultsDesktop(value)
	})
	provinceSelect.PlaceHolder = "Chọn loại vé số"

	dateSelectDesktop = widget.NewSelect([]string{}, func(value string) {
		showResultsDesktop(value)
	})
	dateSelectDesktop.PlaceHolder = "Chọn ngày"

	resultsLabelDesktop = widget.NewLabel("")
	resultsLabelDesktop.Wrapping = fyne.TextWrapWord

	input := widget.NewEntry()
	input.SetPlaceHolder("Nhập số cần kiểm tra")

	checkBtn := widget.NewButton("Kiểm tra", func() {
		youNum := input.Text
		if selectedDateDesktop == "" {
			dialog.ShowInformation("Thông báo", "Hãy chọn ngày trước", w)
			return
		}
		giai, num := utils.CheckWinningNumber(parsedResultsDesktop, selectedDateDesktop, youNum)
		if giai != "" {
			dialog.ShowInformation("Kết quả",
				fmt.Sprintf("Số %s trúng giải %s: %s", youNum, giai, num), w)
		} else {
			dialog.ShowInformation("Kết quả", "Không trúng!", w)
		}
	})

	statusDesktop = widget.NewLabel("")

	left := container.NewVBox(
		container.NewCenter(banner),
		provinceSelect,
		dateSelectDesktop,
		input,
		checkBtn,
		statusDesktop,
	)

	right := container.NewMax(resultsLabelDesktop)

	content := container.NewHSplit(left, right)
	content.SetOffset(0.35)

	w.SetContent(content)
	w.Resize(fyne.NewSize(900, 600))
}

func fetchResultsDesktop(prov string) {
	url := rss.Sources(prov)
	data, err := rss.Fetch(url)
	fyne.Do(func() {
		if err != nil {
			statusDesktop.SetText("Lỗi fetch RSS: " + err.Error())
			return
		}
		res, err := rss.Parse(data)
		if err != nil {
			statusDesktop.SetText("Lỗi parse RSS: " + err.Error())
			return
		}

		parsedResultsDesktop = res
		dateSelectDesktop.Options = []string{}
		for _, r := range res {
			dateSelectDesktop.Options = append(dateSelectDesktop.Options, r.Date)
		}
		dateSelectDesktop.Refresh()
		statusDesktop.SetText("Đã tải xong.")
	})
}

func showResultsDesktop(date string) {
	fyne.Do(func() {
		selectedDateDesktop = date
		found := false
		for _, r := range parsedResultsDesktop {
			if r.Date == date {
				text := fmt.Sprintf("=== %s ===\n", r.Title)
				for _, giai := range configs.Order {
					if so, ok := r.Prizes[giai]; ok {
						text += fmt.Sprintf("Giải %s: %s\n", giai, so)
					}
				}
				resultsLabelDesktop.SetText(text)
				found = true
				break
			}
		}
		if !found {
			resultsLabelDesktop.SetText("!!! Không có kết quả !!!")
		}
	})
}
