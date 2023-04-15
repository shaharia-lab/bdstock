package stock

import (
	"fmt"
	"io/ioutil"
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

func TestStock_GetData(t *testing.T) {
	dseHomePageHTML, _ := readTestDataFile("dse_homepage.html")
	dseCompanyPageHTML, _ := readTestDataFile("dse_company_page.html")

	sc := httpmama.ServerConfig{
		TestEndpoints: []httpmama.TestEndpoint{
			{Path: "/", ResponseString: dseHomePageHTML},
			{Path: "/displayCompany.php", ResponseString: dseCompanyPageHTML},
		},
	}
	ts := httpmama.NewTestServer(sc)
	defer ts.Close()

	os.Setenv("DSE_HOMEPAGE", ts.URL)
	st := NewStock(false)
	data, err := st.GetAllStocks()
	assert.NoError(t, err)

	assert.Equal(t, 330, len(data))

	var testData CompanyStockData
	for _, sd := range data {
		if sd.StockCode == "1JANATAMF" {
			testData = sd
		}
	}

	assert.Equal(t, "1JANATAMF", testData.StockCode)
	assert.Equal(t, "6.10", testData.LastTradingPrice)
	assert.Equal(t, "6.10", testData.ClosingPrice)
	assert.Equal(t, "2:10 PM", testData.LastUpdate)
	assert.Equal(t, "6.10 - 6.10", testData.DaysRange)
	assert.Equal(t, "5.70 - 6.60", testData.WeeksMovingRange)
	assert.Equal(t, "6.10", testData.OpeningPrice)
	assert.Equal(t, "30,022.00", testData.DaysVolume)
	assert.Equal(t, "6.10", testData.AdjustedOpening)
	assert.Equal(t, "14", testData.DaysTrade)
	assert.Equal(t, "6.10", testData.YesterdayClosing)
	assert.Equal(t, "1,768.532", testData.MarketCapitalization)
}
