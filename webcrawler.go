package main

import (
	"fmt"
	"github.com/michaelokarimia/crawlerclient"
)

func main() {
	fmt.Printf("Starting crawlerclient")
	domain := "michaelokarimia.com"
	output := crawlerclient.GetUrls(domain)
	fmt.Printf(output)
}
