package main

import (
	"flag"
	"fmt"
	"strings"
	"github.com/jasonriedel/fetcher/crawler"
	"time"
)

var (
	sites = flag.String("url", "", "Name of sites to crawl comma delimitted")

)
func main() {
	flag.Parse()
	start := time.Now()
	count := 0
	if *sites != "" {
		ch := make(chan string)
		if strings.Contains(*sites, ",") {
			u := strings.Split(*sites, ",")
			for _, cu := range u {
				count++
				go crawler.Crawl(cu, ch) // start goroutine
			}
			for range u {
				fmt.Println(<-ch)
			}
		} else {
			count++
			go crawler.Crawl(*sites, ch) // start goroutine
			fmt.Println(<-ch)
		}
	} else {
		fmt.Println("Please specify urls")
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("Total time: %.2fs - %d site(s)", secs, count)
}
