package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	readability "github.com/go-shiori/go-readability"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		log.Print("need a URL")
	}
	u := args[0]

	article, err := readability.FromURL(u, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to parse %s, %v\n", u, err)
	}

	fmt.Printf("URL     : %s\n", u)
	fmt.Printf("Title   : %s\n", article.Title)
	fmt.Printf("Author  : %s\n", article.Byline)
	fmt.Printf("Length  : %d\n", article.Length)
	fmt.Printf("Excerpt : %s\n", article.Excerpt)
	fmt.Printf("SiteName: %s\n", article.SiteName)
	fmt.Printf("Image   : %s\n", article.Image)
	fmt.Printf("Favicon : %s\n", article.Favicon)
	fmt.Println()
	fmt.Println(article.TextContent)
}
