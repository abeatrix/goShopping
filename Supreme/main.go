package main

import (
	"fmt"
	"encoding/json"
	"github.com/gocolly/colly"
	"os"
	"io/ioutil"
	"log"
)

type Item struct {
	Title	string `json:"title"`
	Link	string `json:"link"`
}

func main() {

	details := make([]Item, 0)

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request){
		fmt.Println("Visiting: ", r.URL.String())
	})

	c.OnHTML(".product-name", func(e *colly.HTMLElement){
		dTitle := e.ChildText(".name-link")
		dLink := "https://www.supremenewyork.com" + e.ChildAttr("a[class=name-link]", "href")

		// fmt.Printf("Link found: %q -> %s\n", e.Text, dLink)

		d := Item{
			Title:	dTitle,
			Link: dLink,
		}

		details = append(details, d)
	})

	c.Visit("https://www.supremenewyork.com/shop/all/accessories")

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	enc.Encode(details)
	toJSON(details)
}

func toJSON(data []Item) {
	newFile, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Stop trying.")
	}

	ioutil.WriteFile("supreme.json", newFile, 0600)
}
