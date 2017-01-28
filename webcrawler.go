package main

import (
	"fmt"
	"github.com/michaelokarimia/crawlerclient"
)

func main() {
	fmt.Printf("Starting crawlerclient")
	domain := "michaelokarimia.com"
	output := crawlerclient.Crawlerclient(domain)
	fmt.Printf(output)
}
