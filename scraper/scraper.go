package scraper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

var C *colly.Collector

func Scraper() error {
	isins := []string{"IE00B1FZS574", "IE00BKX55T58", "IE00B3VVMM84"}
	etfInfo := EtfInfo{}
	etdInfos := make([]EtfInfo, 1, 2)
	link := "www.trackingdifferences.com"

	C = colly.NewCollector(colly.AllowedDomains("www.trackingdifferences.com", "trackingdifferences.com"))

	C.SetCookies(link, []*http.Cookie{})

	C.OnHTML("h1.page-title", func(h *colly.HTMLElement) {
		etfInfo.Title = h.Text
		fmt.Println(h.Text)
	})

	C.OnHTML("div.descfloat p.desc", func(h *colly.HTMLElement) {
		selection := h.DOM
		//fmt.Println("desc find:%s", selection.Find("span.desctitle").Text())
		childNodes := (selection.Children().Nodes)
		if len(childNodes) == 3 {
			description := selection.Find("span.desctitle").Text()
			description = Cleandesk(description)
			value := selection.FindNodes(childNodes[2])
			//fmt.Println("%s, %s \n", description, value)
			switch description {
			case "Replikation":
				etfInfo.Replication = value.Text()
				break
			case "TER":
				etfInfo.TotalExpenseRatio = value.Text()
				break
			case "TD":
				etfInfo.TrackingDifference = value.Text()
				break
			case "Volumen":
				etfInfo.Earnings = value.Text()
				break
			case "Land":
				etfInfo.FundSize = value.Text()
				break
			}
		}
	})

	C.OnScraped(func(r *colly.Response) {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		enc.Encode(etfInfo)
		etdInfos = append(etdInfos, etfInfo)
		etfInfo = EtfInfo{}
	})

	C.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting %v", r.URL)
	})

	C.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error while scrapping %s", err.Error())
	})

	for _, ISIN := range isins {
		// Visit the URL and start scraping
		err := C.Visit(scrapeUrl(ISIN))
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func scrapeUrl(s string) string {
	return "https://www.trackingdifferences.com/ETF/ISIN/" + s
}

type EtfInfo struct {
	Title              string
	Replication        string
	Earnings           string
	TotalExpenseRatio  string
	TrackingDifference string
	FundSize           string
}

func Cleandesk(s string) string {
	return strings.TrimSpace(s)
}
