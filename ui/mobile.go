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
	dateSelectMobile    *widget.Select
	resultsLabelMobile  *widget.Label
	statusMobile        *widget.Label
	parsedResultsMobile []rss.Result
	selectedDateMobile  string
)

func BuildMobileUI(w fyne.Window) {
	banner := canvas.NewText("xskt", color.White)
	banner.TextStyle = fyne.TextStyle{Bold: true}
	banner.TextSize = 36

	provinceSelect := widget.NewSelect(configs.Provinces, func(value string) {
		statusMobile.SetText("Đang tải dữ liệu...")
		go fetchResultsMobile(value)
	})
	provinceSelect.PlaceHolder = "Chọn loại vé số"

	dateSelectMobile = widget.NewSelect([]string{}, func(value string) {
		showResultsMobile(value)
	})
	dateSelectMobile.PlaceHolder = "Chọn ngày"

	resultsLabelMobile = widget.NewLabel("")
	resultsLabelMobile.Wrapping = fyne.TextWrapWord

	input := widget.NewEntry()
	input.SetPlaceHolder("Nhập số cần kiểm tra")

	checkBtn := widget.NewButton("Kiểm tra", func() {
		youNum := input.Text
		if selectedDateMobile == "" {
			dialog.ShowInformation("Thông báo", "Hãy chọn ngày trước", w)
			return
		}
		giai, num := utils.CheckWinningNumber(parsedResultsMobile, selectedDateMobile, youNum)
		if giai != "" {
			dialog.ShowInformation("Kết quả",
				fmt.Sprintf("Số %s trúng giải %s: %s", youNum, giai, num), w)
		} else {
			dialog.ShowInformation("Kết quả", "Không trúng!", w)
		}
	})

	statusMobile = widget.NewLabel("")

	content := container.NewVBox(
		container.NewCenter(banner),
		provinceSelect,
		dateSelectMobile,
		resultsLabelMobile,
		input,
		checkBtn,
		statusMobile,
	)

	w.SetContent(container.NewScroll(content))
	w.Resize(fyne.NewSize(400, 700))
}

func fetchResultsMobile(prov string) {
	url := rss.Sources(prov)
	data, err := rss.Fetch(url)
	fyne.Do(func() {
		if err != nil {
			statusMobile.SetText("Lỗi fetch RSS: " + err.Error())
			return
		}
		res, err := rss.Parse(data)
		if err != nil {
			statusMobile.SetText("Lỗi parse RSS: " + err.Error())
			return
		}

		parsedResultsMobile = res
		dateSelectMobile.Options = []string{}
		for _, r := range res {
			dateSelectMobile.Options = append(dateSelectMobile.Options, r.Date)
		}
		dateSelectMobile.Refresh()
		statusMobile.SetText("Đã tải xong.")
	})
}

func showResultsMobile(date string) {
	fyne.Do(func() {
		selectedDateMobile = date
		found := false
		for _, r := range parsedResultsMobile {
			if r.Date == date {
				text := fmt.Sprintf("=== %s ===\n", r.Title)
				for _, giai := range configs.Order {
					if so, ok := r.Prizes[giai]; ok {
						text += fmt.Sprintf("Giải %s: %s\n", giai, so)
					}
				}
				resultsLabelMobile.SetText(text)
				found = true
				break
			}
		}
		if !found {
			resultsLabelMobile.SetText("!!! Không có kết quả !!!")
		}
	})
}
