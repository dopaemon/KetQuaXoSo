package rss

import (
	"encoding/xml"
	"html"
	"regexp"
	"strings"
	"time"
)

type Result struct {
	Province string
	Date     string
	Title    string
	Prizes   map[string][]string
}

type rssFeed struct {
	Channel struct {
		Title string `xml:"title"`
		Items []struct {
			Title       string `xml:"title"`
			Description string `xml:"description"`
			PubDate     string `xml:"pubDate"`
			Link        string `xml:"link"`
		} `xml:"item"`
	} `xml:"channel"`
}

func Parse(data []byte) ([]Result, error) {
	var feed rssFeed
	if err := xml.Unmarshal(data, &feed); err != nil {
		return nil, err
	}

	var results []Result
	for _, it := range feed.Channel.Items {
		desc := normalizeDescription(it.Description)
		prizes := parsePrizes(desc)

		date := normalizeDate(it.PubDate)
		title := it.Title

		prov := detectProvinceFromItemTitle(it.Title)
		if prov == "" {
			prov = cleanProvinceFromChannel(feed.Channel.Title)
		}

		results = append(results, Result{
			Province: prov,
			Date:     date,
			Title:    title,
			Prizes:   prizes,
		})
	}
	return results, nil
}

func normalizeDescription(s string) string {
	s = html.UnescapeString(s)
	s = strings.ReplaceAll(s, "<br>", " ")
	s = strings.ReplaceAll(s, "<br/>", " ")
	s = strings.ReplaceAll(s, "<br />", " ")

	spaceRe := regexp.MustCompile(`\s+`)
	s = spaceRe.ReplaceAllString(s, " ")

	re78 := regexp.MustCompile(`7:\s*([0-9]{3})\s*(?:-?\s*)8:\s*([0-9]{2})`)
	s = re78.ReplaceAllString(s, "7: $1 8: $2")

	labelRe := regexp.MustCompile(`\s*(ĐB|[1-8])\s*:`)
	s = labelRe.ReplaceAllString(s, "\n$1:")

	return strings.TrimSpace(s)
}

func parsePrizes(s string) map[string][]string {
	out := make(map[string][]string)

	lineRe := regexp.MustCompile(`(?m)^(ĐB|[1-8]):\s*(.+)$`)
	for _, m := range lineRe.FindAllStringSubmatch(s, -1) {
		label := m[1]
		val := m[2]

		splits := strings.FieldsFunc(val, func(r rune) bool { return r == '-' || r == ' ' })
		for _, v := range splits {
			v = strings.TrimSpace(v)
			if v != "" {
				out[label] = append(out[label], v)
			}
		}
	}
	return out
}

func detectProvinceFromItemTitle(title string) string {
	up := strings.ToUpper(title)
	start := strings.Index(up, "XỔ SỐ ")
	if start < 0 {
		return ""
	}
	start += len("XỔ SỐ ")
	end := strings.Index(up[start:], " NGÀY")
	if end < 0 {
		return ""
	}
	raw := strings.TrimSpace(title[start : start+end])
	if raw == "" {
		return ""
	}
	return titleCaseVN(raw)
}

func cleanProvinceFromChannel(chTitle string) string {
	prefix := "Kết quả xổ số "
	if strings.HasPrefix(chTitle, prefix) {
		return titleCaseVN(strings.TrimSpace(chTitle[len(prefix):]))
	}
	return titleCaseVN(strings.TrimSpace(chTitle))
}

func normalizeDate(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}
	layouts := []string{
		time.RFC1123Z,
		time.RFC1123,
		"02/01/2006",
	}
	for _, ly := range layouts {
		if t, err := time.Parse(ly, s); err == nil {
			return t.Format("2006-01-02")
		}
	}
	return s
}

func titleCaseVN(s string) string {
	parts := strings.Fields(strings.ToLower(s))
	for i, p := range parts {
		if len(p) > 0 {
			parts[i] = strings.ToUpper(p[:1]) + p[1:]
		}
	}
	return strings.Join(parts, " ")
}
