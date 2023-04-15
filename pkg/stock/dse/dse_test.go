package dse

import (
	"fmt"
	"github.com/shahariaazam/bdstock/pkg/stock"
	"io/ioutil"
	"net/url"
	"os"
	"testing"

	"github.com/shahariaazam/httpmama"
	"github.com/stretchr/testify/assert"
)

func readTestDataFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(fmt.Sprintf("testdata/%s", filename))
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func TestStock_GetStockData(t *testing.T) {
	tests := []struct {
		name                string
		stockCode           string
		expectedCompanyData stock.CompanyStockData
	}{
		{
			name:      "return all data successfully",
			stockCode: "1JANATAMF",
			expectedCompanyData: stock.CompanyStockData{
				StockCode:            "1JANATAMF",
				LastTradingPrice:     "6.10",
				ClosingPrice:         "6.10",
				LastUpdate:           "2:10 PM",
				DaysRange:            "6.10 - 6.10",
				WeeksMovingRange:     "5.70 - 6.60",
				OpeningPrice:         "6.10",
				DaysVolume:           "30,022.00",
				AdjustedOpening:      "6.10",
				DaysTrade:            "14",
				YesterdayClosing:     "6.10",
				MarketCapitalization: "1,768.532",
			},
		},
	}

	dseCompanyPageHTML, _ := readTestDataFile("dse_company_page.html")

	sc := httpmama.ServerConfig{
		TestEndpoints: []httpmama.TestEndpoint{
			{Path: "/displayCompany.php", ResponseString: dseCompanyPageHTML, QueryParams: url.Values{"name": []string{"1JANATAMF"}}},
		},
	}
	ts := httpmama.NewTestServer(sc)
	defer ts.Close()

	os.Setenv("DSE_HOMEPAGE", ts.URL)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			st := DSE{}
			data, err := st.GetStockData(tc.stockCode)
			assert.NoError(t, err)

			assert.Equal(t, tc.expectedCompanyData, data)
		})
	}
}
