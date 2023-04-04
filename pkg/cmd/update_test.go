package cmd

import (
	"bytes"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func executeCommand(cmd *cobra.Command, args ...string) (string, error) {
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs(args)

	err := cmd.Execute()

	return buf.String(), err
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

func TestGreetCommand(t *testing.T) {
	ts := httptest.NewServer(getHandler())
	defer ts.Close()

	os.Setenv("DSE_ENDPOINT", ts.URL)

	t.Run("Greet with no name", func(t *testing.T) {
		updateCmd := NewUpdateCommand()
		err := updateCmd.Execute()

		assert.NoError(t, err)
	})
}

func readTestDataFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(fmt.Sprintf("testdata/%s", filename))
	if err != nil {
		return "", err
	}

	return string(data), nil
}
