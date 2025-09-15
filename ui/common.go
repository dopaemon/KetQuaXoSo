package ui

import (
	"fmt"
	"net/url"

	"github.com/dopaemon/KetQuaXoSo/internal/configs"
	"github.com/dopaemon/KetQuaXoSo/internal/rss"
	"github.com/dopaemon/KetQuaXoSo/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type UIState struct {
	DateSelect    *widget.Select
	ResultsLabel  *widget.Label
	Status        *widget.Label
	LinkRSS       *widget.Hyperlink
	ParsedResults []rss.Result
	SelectedDate  string
}

func FetchResults(prov string, ui *UIState) {
	ui.Status.Wrapping = fyne.TextWrapWord

	urlrss, code := rss.Sources(prov)
	data, err := rss.Fetch(urlrss)
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
		u, _ := url.Parse("https://xskt.com.vn/xs" + code)
		ui.LinkRSS.SetText("Xem Thêm ...")
		ui.LinkRSS.SetURL(u)
		ui.LinkRSS.Refresh()
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
						text += fmt.Sprintf(" - Giải %s: %s\n", giai, so)
					}
				}
				text += "\nChúc bạn may mắn !!!"
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

func IsSixDigitNumber(s string) bool {
	if len(s) < 5 || len(s) > 6 {
		return false
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
