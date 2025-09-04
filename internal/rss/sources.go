package rss

import (
	"fmt"
	"strings"
)

func Sources(tinh string) string {
	s := tinh
	parts := strings.Split(s, "-")
	var ab string
	for _, p := range parts {
		if len(p) > 0 {
			ab += string(p[0])
		}
	}
	fmt.Println(ab)
	return "https://xskt.com.vn/rss-feed/" + tinh + "-xs" + ab + ".rss"
}
