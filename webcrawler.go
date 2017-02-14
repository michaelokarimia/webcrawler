package main

import (
	"fmt"
	"os"

	"github.com/michaelokarimia/crawlerclient"
)

func main() {
	fmt.Printf("Starting crawlerclient")

	if len(os.Args) == 2 {
		domain := os.Args[1]

		fmt.Printf("\nCrawling domain: %s", domain)

		output := crawlerclient.GetUrls(domain)
		fmt.Printf("\n" + output + "\n")
	} else {
		fmt.Printf("\nUsage: argument must be a valid domain, i.e example.com \n")
	}
}
