package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-shiori/dom"
	"github.com/koron-go/janorm"
	"github.com/koron-go/ngram"
	"golang.org/x/net/html"
)

type entry struct {
	name string
	text string
}

func loadArticles(name string) ([]entry, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var entries []entry
	r := bufio.NewReader(f)
	n := 0
	for {
		l, err := r.ReadString('\n')
		n++
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}
		items := strings.SplitN(strings.TrimSpace(l), "\t", 4)
		if len(items) != 3 && len(items) != 2 {
			return nil, fmt.Errorf("unexpected count of items at line %d: want=2or3, got=%d", n, len(items))
		}
		if items[0] != "OK" {
			log.Printf("skip at line %d: %+v", items[:2])
			continue
		}
		var text string
		if len(items) >= 3 {
			text = items[2]
		}
		entries = append(entries, entry{name: items[1], text: text})
	}
	return entries, nil
}

func readFileAsDOM(name string) (*html.Node, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return dom.Parse(f)
}

func extractMetaDesc(doc *html.Node) string {
	for _, n := range dom.QuerySelectorAll(doc, `meta[name='description']`) {
		return dom.GetAttribute(n, "content")
	}
	return ""
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

func check(e entry) error {
	doc, err := readFileAsDOM(filepath.Join("..", "dataset", e.name))
	if err != nil {
		return err
	}
	descStr := regulateText(extractMetaDesc(doc))
	descIdx := ngram.New(2, descStr)
	textStr := regulateText(e.text)
	textIdx := ngram.New(2, textStr)
	f := calcFrac(descIdx, textIdx)
	fmt.Printf("%s\t%f\t%q\t%s\n", e.name, f, descStr, textStr)
	return nil
}

var input string

func main() {
	flag.StringVar(&input, "input", "../java/test5_java_out.txt", "")
	flag.Parse()
	entries, err := loadArticles(input)
	if err != nil {
		log.Fatalf("failed to load article: %s", err)
	}
	for _, e := range entries {
		err := check(e)
		if err != nil {
			log.Fatalf("failed to check: %+v: %s", e, err)
		}
	}
}
