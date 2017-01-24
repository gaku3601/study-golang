package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/text/encoding/japanese"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/transform"
)

func ExampleScrape() {
	// Load the URL
	res, err := http.Get("http://daily.2ch.net/newsplus/subback.html")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	//shift-jis→UTF8
	scanner := transform.NewReader(res.Body, japanese.ShiftJIS.NewDecoder())

	doc, err := goquery.NewDocumentFromReader(scanner)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("#trad a").Each(func(i int, s *goquery.Selection) {
		//band := s.Find("a").Text()    //要素探索する場合、こういう書き方も可能
		title := s.Text()
		fmt.Printf("Thread %d: %s \n", i, title)
	})
}

func main() {
	ExampleScrape()
}
