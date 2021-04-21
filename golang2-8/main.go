package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"go.uber.org/zap"
)

var logger *zap.Logger

const numberGoroutines = 100

type Result struct {
	file string
	hash string
}

// This function implements file de-duplication (dedup)
// it uses concurrently recursive walk function
//
// Example to use:
// go run ./main.go -d ./test -r true
func main() {
	InitLogger()
	defer logger.Sync()
	logger = logger.With(zap.String("host", "srv42")).With(zap.Uint64("uid", 100500))

	var (
		dir     string
		isDedup bool
	)

	logger.Info("Parse command line arguments")

	flag.StringVar(&dir, "d", ".", "Path to a dir where we should find duplicates")
	flag.BoolVar(&isDedup, "r", false, "Action to remove duplicates")
	var help = flag.Bool("h", false, "Display this message")
	flag.Parse()
	if *help {
		fmt.Println("duplicates is a command line tool to find duplicate files in a folder")
		fmt.Println("usage: duplicates [options...] path")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if len(flag.Args()) < 1 {
		fmt.Fprintf(os.Stderr, "You have to specify at least a directory to explore ...\n")
		os.Exit(-1)
	}

	logger.Info("run main recursive function with " + strconv.Itoa(numberGoroutines) + "goroutines")
	counter, err := run(dir, numberGoroutines)
	if err != nil {
		logger.Error("failed to run thanks to " + err.Error())
		os.Exit(1)
	}

	logger.Info("remove duplicates " + strconv.FormatBool(isDedup) + " if necessary and show the outcome")
	for hash, files := range counter {
		if len(files) > 1 {
			logger.Info("remove duplicates " + strconv.FormatBool(isDedup) + " if necessary and show the outcome")

			fmt.Printf("Found %d duplicates for %v: \n", len(files), hash)
			for i, f := range files {
				// remove elements other than a first one
				if i > 0 && isDedup {
					err = os.Remove(f)
					if err != nil {
						fmt.Printf("can not remove file %s: %s", f, err)
					}
				} else {
					fmt.Println("-> ", f)
				}
			}
		}
	}
}

func InitLogger() {
	logger, _ = zap.NewProduction()
}

func run(dir string, workers int) (map[string][]string, error) {

	input := make(chan string)
	go search(dir, input)

	counter := make(map[string][]string)
	results := make(chan *Result)
	go func() {
		for r := range results {
			counter[r.hash] = append(counter[r.hash], r.file)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go startWorker(input, results, &wg)
	}
	wg.Wait()
	close(results)

	return counter, nil
}

func search(dir string, input chan<- string) {
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logger.Error("file corrupted" + err.Error())
		} else if info.Mode().IsRegular() {
			input <- path
		}
		return nil
	})
	close(input)
}

func startWorker(input <-chan string, results chan<- *Result, wg *sync.WaitGroup) {
	logger.Info("")
	for file := range input {
		fs, err := os.Stat(file)
		if err != nil {
			logger.Error("file corrupted" + err.Error())
			continue
		}
		results <- &Result{
			file: file,
			hash: fmt.Sprintf("%s_%d", fs.Name(), fs.Size()), // this is kind of hash that represents a file
		}
	}
	wg.Done()
}
