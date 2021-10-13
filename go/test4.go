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
		content := regulateText(a.TextContent)
		cIndex := ngram.New(2, content)
		frac := calcFrac(tIndex, cIndex)
		fmt.Printf("OK\t%s\t%f\t%s\t%s\n", u, frac, title, content)
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
