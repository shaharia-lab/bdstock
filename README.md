<h1 align="center">BD Stock Market</h1>
<p align="center">Collect stock price information for Bangladeshi Stock Exchange</p>

<p align="center">
  <a href="https://github.com/shahariaazam/bd-stock-market/actions/workflows/CI.yaml"><img src="https://github.com/shahariaazam/bd-stock-market/actions/workflows/CI.yaml/badge.svg" height="20"/></a>
  <a href="https://codecov.io/gh/shahariaazam/bd-stock-market"><img src="https://codecov.io/gh/shahariaazam/bd-stock-market/branch/master/graph/badge.svg?token=NKTKQ45HDN" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_bd-stock-market"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_bd-stock-market&metric=reliability_rating" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_bd-stock-market"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_bd-stock-market&metric=vulnerabilities" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_bd-stock-market"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_bd-stock-market&metric=security_rating" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_bd-stock-market"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_bd-stock-market&metric=sqale_rating" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_bd-stock-market"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_bd-stock-market&metric=code_smells" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_bd-stock-market"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_bd-stock-market&metric=ncloc" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_bd-stock-market"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_bd-stock-market&metric=alert_status" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_bd-stock-market"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_bd-stock-market&metric=duplicated_lines_density" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_bd-stock-market"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_bd-stock-market&metric=bugs" height="20"/></a>
  <a href="https://sonarcloud.io/summary/new_code?id=shahariaazam_bd-stock-market"><img src="https://sonarcloud.io/api/project_badges/measure?project=shahariaazam_bd-stock-market&metric=sqale_index" height="20"/></a>
</p><br/><br/>

<p align="center">
  <a href="https://github.com/shahariaazam/bd-stock-market"><img src="https://user-images.githubusercontent.com/1095008/230057933-5c4659da-1383-4f99-914d-5bf56c8892fe.png" width="100%"/></a>
</p><br/>


## 🤔  What is BDStock?

**BD Stock Market** a.k.a **bdstock** is a command-line tool that collects and provide stock price information for
Bangladeshi stock exchange market. Currently, it provides the stock information only for [Dhaka Stock Exchange](https://www.dsebd.org). 

## Usage

Download the latest release from [GitHub](https://github.com/shahariaazam/bd-stock-market/releases). And run the
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

Use " [command] --help" for more information about a command
```

To get the latest stock price, please run `update` command

```shell
➜ bdstock update
```

And it will output the stock information in JSON format.

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

Contributions are welcome! Please follow the guidelines outlined in the [CONTRIBUTING](https://github.com/shahariaazam/bd-stock-market/blob/master/CONTRIBUTING.md) file.

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/shahariaazam/bd-stock-market/blob/master/LICENSE) file for details.