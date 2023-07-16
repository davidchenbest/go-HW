package main

import "example.com/hw/scrape"

func main() {
	url := "https://www.amazon.com/gp/product/B08B34HWKV?smid=A2L77EE7U53NWQ"
	scrape.Perform(url)
}

// go run . -rod="show,slow=1s,trace"
