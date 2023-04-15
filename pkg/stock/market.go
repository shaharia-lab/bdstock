package stock

type Market interface {
	GetAllStocks() ([]CompanyStockData, error)
	GetStockInBatches(stockCodes []string, batchSize int) ([]CompanyStockData, []error)
	GetStockInfo(stockCode string) (CompanyStockData, error)
	GetAllStockCodes() ([]string, error)
}
