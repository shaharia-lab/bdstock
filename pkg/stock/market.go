package stock

import (
	"github.com/sirupsen/logrus"
	"sync"
)

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

type Market interface {
	GetAllStocks() ([]CompanyStockData, error)
	GetStockInBatches(batchSize int, stockCodes ...string) ([]CompanyStockData, []error)
	GetSingleStock(stockCode string) (CompanyStockData, error)
	GetStockCodes() ([]string, error)
}

type Data struct {
	logger    *logrus.Logger
	market    Market
	batchSize int
}

func NewStockData(market Market, batchSize int, logger *logrus.Logger) *Data {
	return &Data{
		logger:    logger,
		market:    market,
		batchSize: batchSize,
	}
}

func (d *Data) GetAllStock() ([]CompanyStockData, error) {
	return d.market.GetAllStocks()
}

func (d *Data) GetStocks(stockCodes ...string) ([]CompanyStockData, []error) {
	return d.market.GetStockInBatches(5, stockCodes...)
}

func (d *Data) GetStockInBatches(market Market, stockCodes ...string) ([]CompanyStockData, []error) {
	var stockInfo []CompanyStockData
	var ers []error

	var wg sync.WaitGroup
	var mu sync.Mutex

	i := 1
	perGoroutine := len(stockCodes) / 10
	totalStocks := len(stockCodes)

	for j := 0; j < totalStocks; j += perGoroutine {
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for _, code := range stockCodes[start:end] {

				companyData, err := market.GetSingleStock(code)
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

	return stockInfo, ers
}

// helper function to get minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
