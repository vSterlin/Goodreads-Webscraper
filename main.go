package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

type book struct {
	Title       string
	Author      string
	PublishDate string
}

func main() {

	books := []*book{}
	c := colly.NewCollector()

	c.OnHTML(".leftContainer > .elementList", func(e *colly.HTMLElement) {
		// newUrl := e.Attr("href")
		// e.Request.Visit(newUrl)

		smallText := strings.Split(e.ChildText("span.greyText.smallText"), " ")

		book := &book{
			Title:       e.ChildText(".bookTitle"),
			Author:      e.ChildText("span[itemprop=name]"),
			PublishDate: smallText[len(smallText)-1],
		}

		books = append(books, book)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting " + r.URL.String())
	})

	c.Visit("https://www.goodreads.com/shelf/show/fantasy?page=1")

	c.Wait()

	for _, b := range books {
		fmt.Printf("%+v\n", b)
	}

}
