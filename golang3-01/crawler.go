package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/pkg/errors"
)

type crawlResult struct {
	err error
	msg string
}

type crawler struct {
	sync.Mutex
	visited  map[string]string
	maxDepth int
}

func newCrawler(maxDepth int) *crawler {
	return &crawler{
		visited:  make(map[string]string),
		maxDepth: maxDepth,
	}
}

// scan recursively pages
func (c *crawler) run(ctx context.Context, url string, results chan<- crawlResult, depth int) {
	time.Sleep(2 * time.Second)

	// check if ctx is valid
	select {
	case <-ctx.Done():
		return

	default:
		if depth >= c.maxDepth {
			return
		}

		page, err := parse(url)
		if err != nil {
			// send errors to the channel for futher treatment
			results <- crawlResult{
				err: errors.Wrapf(err, "parse page %s", url),
			}
			return
		}

		title := pageTitle(page)
		links := pageLinks(nil, page)

		// the lock is required since we modify the dict in several goroutines
		c.Lock()
		c.visited[url] = title
		c.Unlock()

		results <- crawlResult{
			err: nil,
			msg: fmt.Sprintf("%s -> %s\n", url, title),
		}

		// look for links recursively
		for link := range links {
			if c.checkVisited(link) {
				continue
			}

			go c.run(ctx, link, results, depth+1)
		}
	}
}

func (c *crawler) checkVisited(url string) bool {
	c.Lock()
	defer c.Unlock()

	_, ok := c.visited[url]
	return ok
}
