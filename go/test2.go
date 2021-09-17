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
	"time"

	readability "github.com/go-shiori/go-readability"
)

func getPlain(u string) (string, error) {
	a, err := readability.FromURL(u, 30*time.Second)
	if err != nil {
		return "", err
	}
	prev := false
	return strings.Map(func(r rune) rune {
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
	}, a.TextContent), nil
}

func fetchPlain(next func() string) {
	for {
		u := next()
		if u == "" {
			return
		}
		text, err := getPlain(u)
		if err != nil {
			fmt.Printf("NG\t%s\t%s\n", u, err)
			continue
		}
		fmt.Printf("OK\t%s\t%s\n", u, text)
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
