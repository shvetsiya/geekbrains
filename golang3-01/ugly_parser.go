package main

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// parse html page
func parse(ctx context.Context, url string) (*html.Node, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		r, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("can't get page")
		}

		b, err := html.Parse(r.Body)
		if err != nil {
			return nil, fmt.Errorf("can't parse page")
		}
		return b, err
	}
}

// pageTitle finds page title
func pageTitle(n *html.Node) string {
	var title string
	if n.Type == html.ElementNode && n.Data == "title" {
		return n.FirstChild.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		title = pageTitle(c)
		if title != "" {
			break
		}
	}
	return title
}

// look for links in the page, use map to exclude duplicates
func pageLinks(links map[string]struct{}, n *html.Node) map[string]struct{} {
	if links == nil {
		links = make(map[string]struct{})
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key != "href" {
				continue
			}

			// it is done for simplicity
			if _, ok := links[a.Val]; !ok && len(a.Val) > 2 && a.Val[:2] == "//" {
				links["http://"+a.Val[2:]] = struct{}{}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = pageLinks(links, c)
	}
	return links
}
