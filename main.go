package main

import (
	"fmt"
	"os"
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
		title := fmt.Sprintf("\"%s\"", e.ChildText(".bookTitle"))

		book := &book{
			Title:       title,
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

	f, err := os.OpenFile("books.csv", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	f.WriteString("Title, Author, Pages\n")
	for _, b := range books {
		s := fmt.Sprintf("%s, %s, %s\n", b.Title, b.Author, b.PublishDate)
		_, err = f.WriteString(s)
		if err != nil {
			fmt.Println(err)
		}
	}

}
