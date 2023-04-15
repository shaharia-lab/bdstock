// Package dse collect stock information
package dse

import (
	"errors"
	"fmt"
	"github.com/shahariaazam/bdstock/pkg/stock"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/kelseyhightower/envconfig"
)

type DSEConfig struct {
	Homepage string `envconfig:"DSE_HOMEPAGE" default:"https://www.dsebd.org/"`
}

type Config struct {
	DSE       DSEConfig
	BatchSize int `envconfig:"DSE_BATCH_SIZE" default:"10"`
}

// Stock processor
type Stock struct {
	verbose bool
	config  Config
}

// NewStock construct new stock processor
func NewStock(verbose bool) *Stock {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil
	}

	return &Stock{verbose: verbose, config: cfg}
}

// GetAllStocks fetch and parse all stocks from Dhaka Stock Exchange
func (s *Stock) GetAllStocks() ([]stock.CompanyStockData, error) {
	stockCodes, err := s.GetAllStockCodes()
	if err != nil {
		return nil, err
	}

	stockInfo, ers := s.GetStockInBatches(stockCodes, s.config.BatchSize)

	if len(ers) == len(stockCodes) {
		return []stock.CompanyStockData{}, fmt.Errorf("failed to get the stock information due to errors")
	}

	return stockInfo, nil
}

// GetStockInBatches fetch and parse stock information in batches with parallel processing
func (s *Stock) GetStockInBatches(stockCodes []string, batchSize int) ([]stock.CompanyStockData, []error) {
	var stockInfo []stock.CompanyStockData
	var ers []error

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
	s.printLog(fmt.Sprintf("started at: %s\n", startTime.Format("2006-01-02 15:04:05")))

	for j := 0; j < totalStocks; j += perGoroutine {
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for _, code := range stockCodes[start:end] {

				s.printLog(fmt.Sprintf("[%d/%d] collecting stock price for %s", i, len(stockCodes), code))

				companyData, err := s.GetStockInfo(code)
				mu.Lock()

				if err != nil {
					ers = append(ers, err)
				} else {
					stockInfo = append(stockInfo, companyData)
				}

				mu.Unlock()
				i++
			}
		}(j, min(j+perGoroutine, len(stockCodes)))
	}
	wg.Wait()

	s.printLog("\nfinished\n")

	elapsedTime := time.Since(startTime)
	s.printLog(fmt.Sprintf("elapsed time: %s seconds", elapsedTime.String()))

	return stockInfo, ers
}

// GetStockInfo fetch and parse stock information for a specific stock or company
func (s *Stock) GetStockInfo(stockCode string) (stock.CompanyStockData, error) {
	companyPageHTML, err := s.getHTML(fmt.Sprintf("/displayCompany.php?name=%s", stockCode))
	if err != nil {
		return stock.CompanyStockData{}, fmt.Errorf("failed to fetch company page. erro: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(companyPageHTML.Body)
	if err != nil {
		return stock.CompanyStockData{}, fmt.Errorf("failed to prepare the parser. error: %w", err)
	}

	return s.ParseCompanyPage(stockCode, doc), nil
}

func (s *Stock) ParseCompanyPage(stockCode string, doc *goquery.Document) stock.CompanyStockData {
	rows := doc.Find("table#company tbody tr")

	return stock.CompanyStockData{
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
}

func (s *Stock) getHTML(url string) (*http.Response, error) {
	resp, err := http.Get(s.config.DSE.Homepage + url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch the page" + url)
	}

	return resp, nil
}

func (s *Stock) printLog(message string) {
	if s.verbose {
		fmt.Println(message)
	}
}

// helper function to get minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (s *Stock) GetAllStockCodes() ([]string, error) {
	dseHomepage, err := s.getHTML("/")
	if err != nil {
		panic(err)
	}

	doc, err := goquery.NewDocumentFromReader(dseHomepage.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare the parser. error: %w", err)
	}

	return s.ParseStockCodes(doc), nil
}

func (s *Stock) ParseStockCodes(doc *goquery.Document) []string {
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
