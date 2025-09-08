package ui

import (
	"fmt"

	"KetQuaXoSo/internal/configs"
	"KetQuaXoSo/internal/rss"
	"KetQuaXoSo/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type UIState struct {
	DateSelect    *widget.Select
	ResultsLabel  *widget.Label
	Status        *widget.Label
	ParsedResults []rss.Result
	SelectedDate  string
}

func FetchResults(prov string, ui *UIState) {
	url := rss.Sources(prov)
	data, err := rss.Fetch(url)
	fyne.Do(func() {
		if err != nil {
			ui.Status.SetText("Lỗi fetch RSS: " + err.Error())
			return
		}
		res, err := rss.Parse(data)
		if err != nil {
			ui.Status.SetText("Lỗi parse RSS: " + err.Error())
			return
		}

		ui.ParsedResults = res
		ui.DateSelect.Options = []string{}
		for _, r := range res {
			ui.DateSelect.Options = append(ui.DateSelect.Options, r.Date)
		}
		ui.DateSelect.Refresh()
		ui.Status.SetText("Đã tải xong.")
	})
}

func ShowResults(date string, ui *UIState) {
	fyne.Do(func() {
		ui.SelectedDate = date
		found := false
		for _, r := range ui.ParsedResults {
			if r.Date == date {
				text := fmt.Sprintf("%s\n\n", r.Title)
				text += fmt.Sprintf("Danh sách giải: \n")
				for _, giai := range configs.Order {
					if so, ok := r.Prizes[giai]; ok {
						text += fmt.Sprintf(" - Giải: %s: %s\n", giai, so)
					}
				}
				text += "\n\nChúc bạn may mắn !!!"
				ui.ResultsLabel.SetText(text)
				found = true
				break
			}
		}
		if !found {
			ui.ResultsLabel.SetText("!!! Không có kết quả cho ngày này !!!")
		}
	})
}

func CheckNumber(input string, ui *UIState, w fyne.Window) {
	if ui.SelectedDate == "" {
		dialog.ShowInformation("Thông báo", "Hãy chọn ngày trước", w)
		return
	}
	giai, num := utils.CheckWinningNumber(ui.ParsedResults, ui.SelectedDate, input)
	if giai != "" {
		dialog.ShowInformation("Kết quả",
			fmt.Sprintf("Số %s trúng giải %s: %s", input, giai, num), w)
	} else {
		dialog.ShowInformation("Kết quả", "Không trúng!", w)
	}
}
