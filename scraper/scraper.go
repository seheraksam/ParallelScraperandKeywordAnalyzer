package scraper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

func Scraper() error {
	isins := []string{"IE00B1FZS574", "IE00BKX55T58", "IE00B3VVMM84"}
	var wg sync.WaitGroup
	var mu sync.Mutex
	etfInfos := []EtfInfo{}

	for _, ISIN := range isins {
		wg.Add(1)

		go func(isin string) {
			defer wg.Done()

			c := colly.NewCollector(colly.AllowedDomains("www.trackingdifferences.com", "trackingdifferences.com"))

			c.SetCookies("www.trackingdifferences.com", []*http.Cookie{})

			etfInfo := EtfInfo{}

			c.OnHTML("h1.page-title", func(h *colly.HTMLElement) {
				etfInfo.Title = h.Text
			})

			c.OnHTML("div.descfloat p.desc", func(h *colly.HTMLElement) {
				selection := h.DOM
				childNodes := selection.Children().Nodes
				if len(childNodes) == 3 {
					description := Cleandesk(selection.Find("span.desctitle").Text())
					value := selection.FindNodes(childNodes[2]).Text()

					switch description {
					case "Replikation":
						etfInfo.Replication = value
					case "TER":
						etfInfo.TotalExpenseRatio = value
					case "TD":
						etfInfo.TrackingDifference = value
					case "Volumen":
						etfInfo.Earnings = value
					case "Land":
						etfInfo.FundSize = value
					}
				}
			})

			c.OnScraped(func(r *colly.Response) {
				mu.Lock()
				etfInfos = append(etfInfos, etfInfo)
				mu.Unlock()
			})

			c.OnError(func(r *colly.Response, err error) {
				log.Printf("Error while scraping %s: %v", isin, err)
			})

			fmt.Printf("Visiting %s\n", scrapeUrl(isin))

			err := c.Visit(scrapeUrl(isin))
			if err != nil {
				log.Println("Visit error:", err)
			}
		}(ISIN) // ISIN değişkenini parametre olarak geçiriyoruz
	}

	wg.Wait() // Tüm goroutinelerin bitmesini bekle

	// Sonuçları JSON olarak yazdır
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(etfInfos)

	return nil
}

func scrapeUrl(s string) string {
	return "https://www.trackingdifferences.com/ETF/ISIN/" + s
}

type EtfInfo struct {
	Title              string `json:"title"`
	Replication        string `json:"replication"`
	Earnings           string `json:"earnings"`
	TotalExpenseRatio  string `json:"total_expense_ratio"`
	TrackingDifference string `json:"tracking_difference"`
	FundSize           string `json:"fund_size"`
}

func Cleandesk(s string) string {
	return strings.TrimSpace(s)
}
