package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/gocolly/colly"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	pUrl := flag.String("url", "", "URL to be processed")
	flag.Parse()
	url := *pUrl
	if url == "" {
		fmt.Fprintf(os.Stderr, "Error: empty URL!\n")
		return
	}

	c := colly.NewCollector(colly.AllowedDomains(strings.Split(*pUrl, "//")[1]))

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "1 Mozilla/5.0 (iPad; CPU OS 12_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148")
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		filename := path.Base(r.Request.URL.String())
		url := r.Ctx.Get("url")

		fmt.Println("Checking if " + filename + " exists ...")
		if _, err := os.Stat(filename); !os.IsNotExist(err) {
			fmt.Println(filename + " already exists!")
			return
		} else {
			f, err := os.Create(filename)
			if err != nil {
				log.Fatal(err)
				return
			}
			defer f.Close()

			if err = r.Save(filename); err != nil {
				log.Fatal(err)
			} else {

				fmt.Println("Downloading ", url, " to ", filename)
				fmt.Printf("filename:%s url:%s", filename, url)
				fmt.Println(filename + " saved!")
				log.Printf("[R] %#v", r.Headers)
			}

		}

	})

	// Start scraping
	c.Visit(*pUrl)
}
