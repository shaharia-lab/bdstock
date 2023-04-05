// Package stock collect stock information
package stock

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const batchSize = 10

// CompanyStockData store stock data
type CompanyStockData struct {
	StockCode            string `json:"stock_code"`
	LastTradingPrice     string `json:"last_trading_price"`
	ClosingPrice         string `json:"closing_price"`
	LastUpdate           string `json:"last_update"`
	DaysRange            string `json:"days_range"`
	WeeksMovingRange     string `json:"weeks_moving_range"`
	OpeningPrice         string `json:"opening_price"`
	DaysVolume           string `json:"days_volume"`
	AdjustedOpening      string `json:"adjusted_opening"`
	DaysTrade            string `json:"days_trade"`
	YesterdayClosing     string `json:"yesterday_closing"`
	MarketCapitalization string `json:"market_capitalization"`
}

// Stock processor
type Stock struct {
	verbose     bool
	dseEndpoint string
}

// NewStock construct new stock processor
func NewStock(dseEndpoint string, verbose bool) *Stock {
	return &Stock{verbose: verbose, dseEndpoint: dseEndpoint}
}

// GetData fetch data from URL
func (s *Stock) GetData() ([]CompanyStockData, error) {
	stockCodes := s.getAllStockCodes()

	stockInfo := s.collectStockPriceInBatch(stockCodes, batchSize)

	return stockInfo, nil
}

func (s *Stock) collectStockPriceInBatch(stockCodes []string, batchSize int) []CompanyStockData {
	var stockInfo []CompanyStockData
	var wg sync.WaitGroup
	var mu sync.Mutex

	i := 1
	perGoroutine := len(stockCodes) / batchSize
	totalStocks := len(stockCodes)

	s.printLog("== Bangladesh Stock Market ==\n")
	s.printLog("started collecting..")
	s.printLog(fmt.Sprintf("per batch: %d", batchSize))
	s.printLog(fmt.Sprintf("total stocks: %d", totalStocks))

	startTime := time.Now()
	s.printLog(fmt.Sprintf("started at: %s\n", startTime.Format("05.04.2022 11:23:00")))

	for j := 0; j < totalStocks; j += perGoroutine {
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for _, code := range stockCodes[start:end] {

				s.printLog(fmt.Sprintf("[%d/%d] collecting stock price for %s", i, len(stockCodes), code))

				companyData := s.getStockInformation(code)
				mu.Lock()
				stockInfo = append(stockInfo, companyData)
				mu.Unlock()
				i++
			}
		}(j, min(j+perGoroutine, len(stockCodes)))
	}
	wg.Wait()

	s.printLog("\nfinished\n")

	elapsedTime := time.Since(startTime)
	s.printLog(fmt.Sprintf("elapsed time: %s seconds", elapsedTime.String()))

	return stockInfo
}

// helper function to get minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (s *Stock) getAllStockCodes() []string {
	dseHomepage, err := s.getHTML("/")
	if err != nil {
		panic(err)
	}

	return s.parseStockCodes(dseHomepage)
}

func (s *Stock) parseStockCodes(doc *goquery.Document) []string {
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

func (s *Stock) getStockInformation(stockCode string) CompanyStockData {
	companySpecificPage, err := s.getHTML(fmt.Sprintf("/displayCompany.php?name=%s", stockCode))
	if err != nil {
		panic(err)
	}

	return s.parseCompanyPageData(stockCode, companySpecificPage)
}

func (s *Stock) parseCompanyPageData(stockCode string, doc *goquery.Document) CompanyStockData {
	// Parse table rows using goquery
	rows := doc.Find("table#company tbody tr")

	companyData := CompanyStockData{
		StockCode:            stockCode,
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

func (s *Stock) getHTML(url string) (*goquery.Document, error) {
	resp, err := http.Get(s.dseEndpoint + url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch the page" + url)
	}

	// Parse the response body using goquery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func (s *Stock) printLog(message string) {
	if s.verbose {
		fmt.Println(message)
	}
}
