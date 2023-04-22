// Package dse collect stock information
package dse

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/shahariaazam/bdstock/pkg/stock"
	"github.com/sirupsen/logrus"
)

// DSE processor
type DSE struct {
	Logger *logrus.Logger
}

func (s *DSE) resolveHomepage() string {
	if found, ok := os.LookupEnv("DSE_HOMEPAGE"); ok {
		return found
	}

	return "https://www.dsebd.org/"
}

// GetStockData fetch and parse stock information for a specific stock or company
func (s *DSE) GetStockData(stockCode string) (stock.CompanyStockData, error) {
	companyPageHTML, err := s.getHTML(fmt.Sprintf("/displayCompany.php?name=%s", stockCode))
	if err != nil {
		return stock.CompanyStockData{}, fmt.Errorf("failed to fetch company page. erro: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(companyPageHTML.Body)
	if err != nil {
		return stock.CompanyStockData{}, fmt.Errorf("failed to prepare the parser. error: %w", err)
	}

	return s.parseCompanyPage(stockCode, doc), nil
}

func (s *DSE) GetStockCodes() ([]string, error) {
	dseHomepage, err := s.getHTML("/")
	if err != nil {
		panic(err)
	}

	doc, err := goquery.NewDocumentFromReader(dseHomepage.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare the parser. error: %w", err)
	}

	return s.parseStockCodes(doc), nil
}

func (s *DSE) parseCompanyPage(stockCode string, doc *goquery.Document) stock.CompanyStockData {
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

func (s *DSE) getHTML(url string) (*http.Response, error) {
	homepage := s.resolveHomepage()
	resp, err := http.Get(homepage + url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch the page" + url)
	}

	return resp, nil
}

func (s *DSE) parseStockCodes(doc *goquery.Document) []string {
	var stockCodes []string

	doc.Find("a.abhead").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		code := strings.Split(href, "=")[1]

		stockCodes = append(stockCodes, code)
	})

	return stockCodes
}
