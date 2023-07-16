// Command text is a chromedp example demonstrating how to extract text from a
// specific element.
package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
	)
	ctx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	// create context
	ctx, cancel := chromedp.NewContext(
		ctx,
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// run task list
	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://stockx.com/sell/air-jordan-1-high-og-spider-man-across-the-spider-verse`),
		// chromedp.Text(`.Documentation-overview`, &res, chromedp.NodeVisible),
		chromedp.Text(`h2`, &res, chromedp.NodeVisible),
	)
	if err != nil {
		log.Fatal(err)
	}

	// log.Println(strings.TrimSpace(res))
	time.Sleep(time.Hour)
}
