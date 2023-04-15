<h1 align="center">BD Stock Market</h1>
<p align="center">Collect stock price information for Bangladeshi Stock Exchange</p>

<p align="center">
  <a href="https://github.com/shahariaazam/bdstock/actions/workflows/CI.yaml"><img src="https://github.com/shahariaazam/bdstock/actions/workflows/CI.yaml/badge.svg" height="20"/></a>
  <a href="https://codecov.io/gh/shahariaazam/bdstock"><img src="https://codecov.io/gh/shahariaazam/bdstock/branch/master/graph/badge.svg?token=NKTKQ45HDN" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_bdstock"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_bdstock&metric=reliability_rating" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_bdstock"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_bdstock&metric=vulnerabilities" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_bdstock"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_bdstock&metric=security_rating" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_bdstock"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_bdstock&metric=sqale_rating" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_bdstock"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_bdstock&metric=code_smells" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_bdstock"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_bdstock&metric=ncloc" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_bdstock"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_bdstock&metric=alert_status" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_bdstock"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_bdstock&metric=duplicated_lines_density" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_bdstock"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_bdstock&metric=bugs" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_bdstock"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_bdstock&metric=sqale_index" height="20"/></a>
</p><br/><br/>

<p align="center">
  <a href="https://github.com/shahariaazam/bdstock"><img src="https://user-images.githubusercontent.com/1095008/230057933-5c4659da-1383-4f99-914d-5bf56c8892fe.png" width="100%"/></a>
</p><br/>


## 🤔  What is BDStock?

**BD Stock Market** a.k.a **bdstock** is a command-line tool and a library in Go to collect and provide stock price 
information for Bangladeshi stock exchange market. Currently, it provides the stock information only for
[Dhaka Stock Exchange](https://www.dsebd.org).

## Use as a library

### Get stock price for a single Company

```go
package main

import (
	"fmt"
	"github.com/shahariaazam/bdstock/pkg/stock"
)

func main() {
	bds := stock.NewStock("https://www.dsebd.org/", false)
	si := bds.GetStockInformation("1JANATAMF")
	fmt.Printf("closing price for %s is %s\n", si.StockCode, si.ClosingPrice)
}
```

### Get stock prices for multiple Company

Also, if you want to get the stock information for many companies at once in a batch,

```go
bd := stock.NewStock("https://www.dsebd.org/", false)
si := bd.GetDataInBatch([]string{"1JANATAMF", "NAVANAPHAR"}, 20)
for _, s := range si {
    fmt.Printf("closing price for %s is %s\n", s.StockCode, s.ClosingPrice)
}
```

In the above code, it would fetch and parse all the stock price in a batch mode (per batch 20)

## Use as a Command

Download the latest release from [GitHub](https://github.com/shahariaazam/bdstock/releases). And run the
program. Here is the command details.

```shell
➜ bdstock                          
Get the stock price information from Bangladesh Stock market

Usage:
   [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  update      Update stock price information for companies

Flags:
  -h, --help      help for this command
  -v, --version   version for this command

Use " [command] --help" for more information about a command.
```

To get the latest stock price, please run `update` command

```shell
➜ bdstock update --file savedata.json
```

And it will output the stock information in JSON format (if you don't provide `--file` flag). If you add `--file`
flag, the stock information will be saved to the file.

```shell
[
    {
       "StockCode": "USMANIAGL",
       "LastTradingPrice": "58.90",
       "ClosingPrice": "58.90",
       "LastUpdate": "2:10 PM",
       "DaysRange": "58.20 - 60.60",
       "WeeksMovingRange": "48.70 - 86.60",
       "OpeningPrice": "60.60",
       "DaysVolume": "10,501.00",
       "AdjustedOpening": "60.20",
       "DaysTrade": "77",
       "YesterdayClosing": "60.20",
       "MarketCapitalization": "1,048.136"
    },
    ....
    ....
]
```

## Disclaimer

The stock price collector tool is provided for informational purposes only. The tool is designed to collect stock price information as accurately as possible, but we do not guarantee the accuracy, completeness, timeliness, or reliability of the information provided by the tool.

The tool is not intended to provide investment advice, and any decisions made based on the information collected by the tool are made at your own risk. We are not responsible for any trading or investment decisions made using the information collected by the tool.

You should not rely solely on the information collected by the tool for making investment decisions. You should conduct your own research and analysis and seek the advice of qualified professionals before making any investment decisions.

We disclaim all liability for any damages or losses, including direct, indirect, incidental, consequential, or punitive damages, arising from the use of the tool or the information collected by the tool.

## 🤝 Contributing

Contributions are welcome! Please follow the guidelines outlined in the [CONTRIBUTING](https://github.com/shahariaazam/bdstock/blob/master/CONTRIBUTING.md) file.

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/shahariaazam/bdstock/blob/master/LICENSE) file for details.