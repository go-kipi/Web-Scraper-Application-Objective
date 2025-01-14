package scrape

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type ScrapeModel struct {
	Id       *int   `json:"id" gorm:"primary_key"`
	Name     string `json:"name" gorm:"column:name"`
	ImageUrl string `json:"image_url" gorm:"column:image_url"`
}
type ScrapeModels []ScrapeModel

func Run(url string) (ScrapeModels, error) {
	var sm ScrapeModels

	res, err := http.Get(url)
	if err != nil {
		return sm, fmt.Errorf("http.Get - %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return sm, fmt.Errorf("status code error: %d %s %v", res.StatusCode, res.Status, err)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return sm, fmt.Errorf("goquery.NewDocumentFromReader - %v", err)
	}

	doc.Find(".grid-col tr").Each(func(i int, cell *goquery.Selection) {
		name := cell.Find(".cell-name .ent-name").Text()
		img := cell.Find("td .img-fixed")
		pic, _ := img.Attr("src")
		if name != "" && pic != "" {
			sm = append(sm, ScrapeModel{
				Name:     name,
				ImageUrl: pic,
			})
		}

	})

	return sm, nil
}
