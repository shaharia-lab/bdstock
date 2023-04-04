//Package main
package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

// Company store company stock information
type Company struct {
	LastTradingPrice     string
	ClosingPrice         string
	LastUpdate           string
	DaysRange            string
	WeeksMovingRange     string
	OpeningPrice         string
	DaysVolume           string
	AdjustedOpening      string
	DaysTrade            string
	YesterdayClosing     string
	MarketCapitalization string
}

func fetchCompanyPage(url string) (*goquery.Document, error) {
	// Make an HTTP GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse the response body using goquery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func parseCompanyPageData(doc *goquery.Document) Company {
	// Parse table rows using goquery
	rows := doc.Find("table#company tbody tr")

	companyData := Company{
		LastTradingPrice:     rows.Eq(1).Find("td").Eq(1).Text(),
		ClosingPrice:         rows.Eq(1).Find("td").Eq(1).Text(),
		LastUpdate:           rows.Eq(2).Find("td").Eq(0).Text(),
		DaysRange:            rows.Eq(2).Find("td").Eq(1).Text(),
		WeeksMovingRange:     rows.Eq(4).Find("td").Eq(1).Text(),
		OpeningPrice:         rows.Eq(5).Find("td").Eq(0).Text(),
		DaysVolume:           rows.Eq(5).Find("td").Eq(1).Text(),
		AdjustedOpening:      rows.Eq(6).Find("td").Eq(0).Text(),
		DaysTrade:            rows.Eq(6).Find("td").Eq(1).Text(),
		YesterdayClosing:     rows.Eq(7).Find("td").Eq(0).Text(),
		MarketCapitalization: rows.Eq(7).Find("td").Eq(1).Text(),
	}

	return companyData
}

func getAllStockCodes() []string {
	old, err := fetchHomepage("https://www.dsebd.org/")
	if err != nil {
		panic(err)
	}

	return parseStockCodes(old)
}

func getStockInformation(stockCode string) Company {
	page, err := fetchCompanyPage(fmt.Sprintf("https://www.dsebd.org/displayCompany.php?name=%s", stockCode))
	if err != nil {
		panic(err)
	}

	return parseCompanyPageData(page)
}

func fetchHomepage(url string) (*goquery.Document, error) {
	// Make an HTTP GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse the response body using goquery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
func parseStockCodes(doc *goquery.Document) []string {
	var stockCodes []string

	// Find all <a> elements with class "abhead" and extract the company stock code
	doc.Find("a.abhead").Each(func(i int, s *goquery.Selection) {
		// Extract the code from the <a> element's href attribute
		href, _ := s.Attr("href")
		code := strings.Split(href, "=")[1]

		stockCodes = append(stockCodes, code)
	})

	return stockCodes
}

func main() {
	// Fetch all stock codes
	stockCodes := getAllStockCodes()

	// Create a channel to receive the stock information
	stockInfoChan := make(chan Company)

	// Fetch stock information in multiple goroutines
	for _, code := range stockCodes {
		go func(code string) {
			stockInfoChan <- getStockInformation(code)
		}(code)
	}

	// Collect stock information from the channel
	var stockInfo []Company
	for i := 0; i < len(stockCodes); i++ {
		stockInfo = append(stockInfo, <-stockInfoChan)
	}

	// Print the stock information
	for i, code := range stockCodes {
		info := stockInfo[i]
		fmt.Printf("\n\n==%s==\nLast Trading Price: %s\nClosing Price: %s\nLast Update: %s\nDay's Range: %s\nWeeks Moving Range: %s\nOpening Price: %s\nDays Volume: %s\nAdjusted Opening: %s\nDays Trade: %s\nYesterday Closing: %s\nMarket Capitalization: %s\n",
			code, info.LastTradingPrice, info.ClosingPrice, info.LastUpdate, info.DaysRange, info.WeeksMovingRange, info.OpeningPrice, info.DaysVolume, info.AdjustedOpening, info.DaysTrade, info.YesterdayClosing, info.MarketCapitalization)
	}
}
