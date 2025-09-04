package rss

import (
	"net/http"

	"github.com/mmcdole/gofeed"
)

type RssItem struct {
	Title       string
	Description string
	Link        string
	PubDate     string
}

func Fetch(url string) ([]RssItem, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	parser := gofeed.NewParser()
	feed, err := parser.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	items := []RssItem{}
	for _, i := range feed.Items {
		items = append(items, RssItem{
			Title:       i.Title,
			Description: i.Description,
			Link:        i.Link,
			PubDate:     i.Published,
		})
	}
	return items, nil
}
