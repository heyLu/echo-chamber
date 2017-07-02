package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var config struct {
	url       string
	perSecond int
	verbose   bool
}

func init() {
	flag.StringVar(&config.url, "url", "http://localhost:12345", "The url to generate load at")
	flag.IntVar(&config.perSecond, "per-second", 10, "The number of requests to generate per second")
	flag.BoolVar(&config.verbose, "verbose", false, "Enable verbose logging")
}

func main() {
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	for {
		generateLoad(config.perSecond, config.url, config.verbose)

		time.Sleep(60 * time.Second)
	}
}

func generateLoad(perSecond int, url string, verbose bool) {
	for i := 0; i < perSecond*60; i++ {
		go func() {
			d := time.Duration(rand.Intn(60))*time.Second + time.Duration(rand.Intn(1000))*time.Millisecond
			time.Sleep(d)

			start := time.Now()
			resp, err := http.Get(url)
			if err != nil {
				exit("GET", err)
				return
			}
			defer resp.Body.Close()

			io.Copy(ioutil.Discard, resp.Body)

			if verbose {
				fmt.Printf("GET %s = %d (took %s)\n", url, resp.StatusCode, time.Since(start))
			}
		}()
	}
}

func exit(msg string, err error) {
	fmt.Fprintf(os.Stderr, "Error: %s: %s\n", msg, err)
	os.Exit(1)
}
