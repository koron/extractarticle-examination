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
		fmt.Printf("OK\t%s\t%s\t%s\n", u, regulateText(a.Title), regulateText(a.TextContent))
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
