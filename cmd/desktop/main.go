package main

import (
	"fmt"
	"image/color"

	"XoSoToanQuoc/internal/configs"
	"XoSoToanQuoc/internal/rss"
	"XoSoToanQuoc/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var (
	dateSelect    *widget.Select
	resultsText   *widget.Entry
	status        *widget.Label
	parsedResults []rss.Result
	selectedDate  string
)

func main() {
	a := app.NewWithID("com.dopaemon.ketquaxoso")
	w := a.NewWindow("Kết Quả Xổ Số")

	bannerText := "xskt"
	banner := canvas.NewText(bannerText, color.White)
	banner.TextStyle = fyne.TextStyle{Bold: true}
	banner.TextSize = 36

	provinceSelect := widget.NewSelect(configs.Provinces, func(value string) {
		status.SetText("Đang tải dữ liệu...")
		status.Refresh()
		go fetchResults(value)
	})
	provinceSelect.PlaceHolder = "Chọn loại vé số"

	dateSelect = widget.NewSelect([]string{}, func(value string) {
		showResults(value)
	})
	dateSelect.PlaceHolder = "Chọn ngày"

	resultsText = widget.NewMultiLineEntry()
	resultsText.SetText("")
	resultsText.SetMinRowsVisible(20)
	resultsText.Disable()

	input := widget.NewEntry()
	input.SetPlaceHolder("Nhập số của bạn để kiểm tra")

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

	leftPanel := container.NewVBox(
		container.NewCenter(banner),
		provinceSelect,
		dateSelect,
		input,
		checkBtn,
		status,
	)

	rightPanel := container.NewMax(resultsText)

	content := container.NewHSplit(leftPanel, rightPanel)
	content.Offset = 0.35

	w.SetContent(content)
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}

func fetchResults(prov string) {
	url := rss.Sources(prov)
	data, err := rss.Fetch(url)

	fyne.Do(func() {
		if err != nil {
			status.SetText("Lỗi fetch RSS: " + err.Error())
			status.Refresh()
			return
		}

		res, err := rss.Parse(data)
		if err != nil {
			status.SetText("Lỗi parse RSS: " + err.Error())
			status.Refresh()
			return
		}

		parsedResults = res
		dateSelect.Options = []string{}
		for _, r := range res {
			dateSelect.Options = append(dateSelect.Options, r.Date)
		}
		dateSelect.Refresh()

		status.SetText("Đã tải xong.")
		status.Refresh()
	})
}

func showResults(date string) {
	fyne.Do(func() {
		selectedDate = date
		found := false
		for _, r := range parsedResults {
			if r.Date == date {
				text := fmt.Sprintf("=== %s ===\n", r.Title)
				for _, giai := range configs.Order {
					if so, ok := r.Prizes[giai]; ok {
						text += fmt.Sprintf("Giải %s: %s\n", giai, so)
					}
				}
				resultsText.SetText(text)
				resultsText.Refresh()
				found = true
				break
			}
		}
		if !found {
			resultsText.SetText("!!! Không có kết quả cho ngày này !!!")
			resultsText.Refresh()
		}
	})
}
