package stock

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
	GetStockInBatches(stockCodes []string, batchSize int) ([]CompanyStockData, []error)
	GetStockInfo(stockCode string) (CompanyStockData, error)
	GetAllStockCodes() ([]string, error)
}
