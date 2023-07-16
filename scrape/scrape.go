package scrape

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/stealth"
)

func Perform(url string) {

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

	details := [2]string{
		page.MustElement("#productTitle").MustText(),
		page.MustElement(".a-price-whole").MustText() + page.MustElement(".a-price-fraction").MustText(),
	}

	fmt.Println(details)

	// time.Sleep(time.Hour)
}
