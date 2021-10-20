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

	readability "github.com/go-shiori/go-readability"
	"github.com/koron-go/janorm"
	"github.com/koron-go/ngram"
)

func extractArticle(name string) (readability.Article, error) {
	f, err := os.Open(name)
	if err != nil {
		return readability.Article{}, err
	}
	defer f.Close()
	return readability.FromReader(f, nil)
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
func fetchPlain(next func() string) {
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
		content := regulateText(a.TextContent)
		cIndex := ngram.New(2, content)
		frac1 := calcFrac(tIndex, cIndex)
		frac2 := calcFrac(eIndex, cIndex)
		fmt.Printf("OK\t%s\t%f\t%f\t%s\t%s\t%s\n", u, frac1, frac2, title, excerpt, content)
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
