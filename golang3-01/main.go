package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	// max numbers of errers allowed during parsing
	errorsLimit = 100000

	// number of results we want to achieve
	resultsLimit = 10000
)

var (
	// url address (e.g. https://en.wikipedia.org/wiki/Lionel_Messi)
	url string

	// how deep we want to look for (e.g. 10)
	depthLimit int
)

// init function starts first
func init() {
	flag.StringVar(&url, "url", "", "url address")
	flag.IntVar(&depthLimit, "depth", 3, "max depth for run")
	flag.Parse()

	// check necessary condition
	if url == "" {
		log.Print("no url set by flag")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {
	started := time.Now()
	crawler := newCrawler(depthLimit)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	go watchSignals(ctx, cancel, crawler)
	defer cancel()

	// create chan for results
	results := make(chan crawlResult)

	// run goroutine to read from channels
	done := watchCrawler(ctx, results, errorsLimit, resultsLimit)

	// the main logic is here. The inner implementation uses recurcive calls inside other goroutines
	crawler.run(ctx, url, results, 0)

	<-done

	log.Println(time.Since(started))
}

func watchSignals(ctx context.Context, cancel context.CancelFunc, crawler *crawler) {
	osSignalChan := make(chan os.Signal)

	signal.Notify(osSignalChan,
		syscall.SIGINT,
		syscall.SIGTERM)

	depthChan := make(chan os.Signal)
	signal.Notify(depthChan, syscall.SIGUSR1)

	for {
		select {
		case <-ctx.Done():
			log.Printf("context timeout is exceeded ")
			// if the signal received, cancel the context
			cancel()
			return
		case sig := <-osSignalChan:
			log.Printf("got a signal %s", sig.String())
			// if the signal received, cancel the context
			cancel()
			return
		case sig := <-depthChan:
			crawler.maxDepth += 2
			log.Printf("received signal %s to increase max depth to %d", sig.String(), crawler.maxDepth)
		}

		sig := <-osSignalChan
		log.Printf("got signal %q", sig.String())
	}
}

func watchCrawler(ctx context.Context, results <-chan crawlResult, maxErrors, maxResults int) chan struct{} {
	readersDone := make(chan struct{})

	go func() {
		defer close(readersDone)
		for {
			select {
			case <-ctx.Done():
				return

			case result := <-results:
				if result.err != nil {
					maxErrors--
					if maxErrors <= 0 {
						log.Println("max errors exceeded")
						return
					}
					continue
				}

				log.Printf("crawling result: %v", result.msg)
				maxResults--
				if maxResults <= 0 {
					log.Println("got max results")
					return
				}
			}
		}
	}()

	return readersDone
}
