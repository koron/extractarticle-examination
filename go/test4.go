package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/go-shiori/dom"
	readability "github.com/go-shiori/go-readability"
	"github.com/koron-go/janorm"
	"github.com/koron-go/ngram"
	"golang.org/x/net/html"
)

type Article struct {
	readability.Article

	HeadTitle string
	MetaDesc  string
}

func readFileAsDOM(name string) (*html.Node, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return dom.Parse(f)
}

func extractHeadTitle(doc *html.Node) string {
	for _, n := range dom.GetElementsByTagName(doc, "title") {
		return dom.TextContent(n)
	}
	return ""
}

func extractMetaDesc(doc *html.Node) string {
	for _, n := range dom.QuerySelectorAll(doc, `meta[name='description']`) {
		return dom.GetAttribute(n, "content")
	}
	return ""
}

func extractArticle(name string) (*Article, error) {
	doc, err := readFileAsDOM(name)
	p := readability.NewParser()
	a, err := p.ParseDocument(doc, nil)
	if err != nil {
		return nil, err
	}
	title := extractHeadTitle(doc)
	desc := extractMetaDesc(doc)
	return &Article{
		Article:   a,
		HeadTitle: title,
		MetaDesc:  desc,
	}, nil
}

func regulateText(s string) string {
	prev := false
	s = strings.Map(func(r rune) rune {
		if r == '　' {
			r = ' '
		}
		switch r {
		case '\t', '\n', '\r':
			prev = false
			return -1
		case ' ':
			if prev {
				return -1
			}
			prev = true
		default:
			prev = false
		}
		return r
	}, s)
	return janorm.Normalize(s)
}

func calcFrac(base, target ngram.Index) float64 {
	cnt := 0
	for k := range base {
		if _, ok := target[k]; ok {
			cnt++
		}
	}
	return float64(cnt) / float64(len(base))
}

func trendLabel(base, target float64) string {
	if target > base {
		return "↗"
	}
	if target < base {
		return "↘"
	}
	return "→"
}
func fetchPlain(next func() string) {
	for i, s := range []string{
		"R",                   // 1
		"Filepath",            // 2
		"F(Title)",            // 3
		"F(head>title)",       // 4
		"F(Excerpt)",          // 5
		"F(meta/description)", // 6
		"Title",               // 7
		"Excerpt",             // 8
		"head>title",          // 9
		"meta/description",    // 10
		"Body",                // 11
	} {
		if i != 0 {
			fmt.Printf("\t")
		}
		fmt.Printf("%d:%s", i+1, s)
	}
	fmt.Println("")

	for {
		u := next()
		if u == "" {
			return
		}
		a, err := extractArticle(u)
		if err != nil {
			fmt.Printf("NG\t%s\t%s\n", u, err)
			continue
		}
		title := regulateText(a.Title)
		tIndex := ngram.New(2, title)
		excerpt := regulateText(a.Excerpt)
		eIndex := ngram.New(2, excerpt)
		headTitle := regulateText(a.HeadTitle)
		htIndex := ngram.New(2, headTitle)
		metaDesc := regulateText(a.MetaDesc)
		mdIndex := ngram.New(2, metaDesc)
		content := regulateText(a.TextContent)
		cIndex := ngram.New(2, content)

		frac1 := calcFrac(tIndex, cIndex)
		frac2 := calcFrac(eIndex, cIndex)
		frac3 := calcFrac(htIndex, cIndex)
		frac4 := calcFrac(mdIndex, cIndex)

		trend3 := trendLabel(frac1, frac3)
		trend4 := trendLabel(frac2, frac4)

		fmt.Printf("OK\t%s\t%f\t%f%s\t%f\t%f%s\t%s\t%s\t%s\t%s\t%s\n",
			u,             // 2
			frac1,         // 3
			frac3, trend3, // 4
			frac2,         // 5
			frac4, trend4, // 6
			title,     // 7
			excerpt,   // 8
			headTitle, // 9
			metaDesc,  // 10
			content,   // 11
		)
	}
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) > 0 {
		i := 0
		fetchPlain(func() string {
			if i >= len(args) {
				return ""
			}
			curr := args[i]
			i++
			return curr
		})
		return
	}

	r := bufio.NewReader(os.Stdin)
	fetchPlain(func() string {
		l, err := r.ReadString('\n')
		if err != nil {
			if !errors.Is(err, io.EOF) {
				log.Printf("failed to read stdin: %s", err)
			}
			return ""
		}
		return strings.TrimSpace(l)
	})
}
