package rss

import (
	_ "fmt"

	"strings"
	"unicode"
)

func Sources(tinh string) string {
	removeTone := func(s string) string {
		var b strings.Builder
		for _, r := range s {
			switch unicode.ToLower(r) {
			case 'á', 'à', 'ả', 'ã', 'ạ', 'ă', 'ắ', 'ằ', 'ẳ', 'ẵ', 'ặ', 'â', 'ấ', 'ầ', 'ẩ', 'ẫ', 'ậ':
				b.WriteRune('a')
			case 'đ':
				b.WriteRune('d')
			case 'é', 'è', 'ẻ', 'ẽ', 'ẹ', 'ê', 'ế', 'ề', 'ể', 'ễ', 'ệ':
				b.WriteRune('e')
			case 'í', 'ì', 'ỉ', 'ĩ', 'ị':
				b.WriteRune('i')
			case 'ó', 'ò', 'ỏ', 'õ', 'ọ', 'ô', 'ố', 'ồ', 'ổ', 'ỗ', 'ộ', 'ơ', 'ớ', 'ờ', 'ở', 'ỡ', 'ợ':
				b.WriteRune('o')
			case 'ú', 'ù', 'ủ', 'ũ', 'ụ', 'ư', 'ứ', 'ừ', 'ử', 'ữ', 'ự':
				b.WriteRune('u')
			case 'ý', 'ỳ', 'ỷ', 'ỹ', 'ỵ':
				b.WriteRune('y')
			default:
				b.WriteRune(r)
			}
		}
		return b.String()
	}

	tinh = strings.ToLower(removeTone(tinh))
	tinh = strings.ReplaceAll(tinh, " ", "-")

	words := strings.Split(tinh, "-")
	code := ""
	for _, w := range words {
		if len(w) > 0 {
			code += string(w[0])
		}
	}

	switch tinh {
		case "binh-dinh": code = "bdi"
		case "da-nang": code = "dng"
		case "dak-nong": code = "dno"
		case "quang-ngai": code = "qng"
		case "quang-nam": code = "qna"
	}
	source := "https://xskt.com.vn/rss-feed/" + tinh + "-xs" + code + ".rss"

	// fmt.Println(source)

	return source
}
