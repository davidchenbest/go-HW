package main

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/stealth"
)

func main() {
	url := "https://stockx.com/sell/air-jordan-1-high-og-spider-man-across-the-spider-verse"
	browser := rod.New().MustConnect().NoDefaultDevice()

	page := stealth.MustPage(browser)
	page.MustNavigate(url)

	// page := browser.MustPage(url)
	// .MustWindowFullscreen()

	// page.MustElement("#searchInput").MustInput("earth")
	// page.MustElement("#search-form > fieldset > button").MustClick()

	// el := page.MustElement("#mw-content-text > div.mw-parser-output > p:nth-child(7)")
	// fmt.Println(el.MustText())

	// page.MustWaitStable().MustScreenshot("a.png")

	el := page.MustElement("h1")
	fmt.Println(el.MustText())

	time.Sleep(time.Hour)
}

// go run . -rod="show,slow=1s,trace"
