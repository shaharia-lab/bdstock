package stock

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func readTestDataFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(fmt.Sprintf("testdata/%s", filename))
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func getHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fileContent, _ := readTestDataFile("dse_homepage.html")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, fileContent)
	})

	mux.HandleFunc("/displayCompany.php", func(w http.ResponseWriter, r *http.Request) {
		fileContent, _ := readTestDataFile("dse_company_page.html")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, fileContent)
	})

	return mux
}

func TestStock_GetData(t *testing.T) {

	ts := httptest.NewServer(getHandler())
	defer ts.Close()

	st := NewStock(ts.URL, false)
	data, err := st.GetData()
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
