# Go-TuShare

Golang版本的TuShare数据接口SDK，支持股票、指数等金融数据查询和分析。

## 安装

```bash
go get github.com/Premium-Platform/go-tushare
```

## 使用方法

### 基本查询

```go
package main

import (
	"fmt"
	"github.com/Premium-Platform/go-tushare/client"
	"os"
)

func main() {
	// 初始化客户端
	cli := client.New("your_token_here")
	// 或者从环境变量获取
	// cli := client.New(os.Getenv("TUSHARE_TOKEN"))
	
	// 查询日线行情数据
	df, err := cli.Query("daily", map[string]interface{}{
		"ts_code":    "000001.SZ",
		"start_date": "20220101",
		"end_date":   "20220110",
	}, []string{})
	
	if err != nil {
		fmt.Printf("获取数据失败: %v\n", err)
		return
	}
	
	// 使用数据
	for _, row := range df.Rows {
		fmt.Printf("日期: %s, 开盘价: %v, 收盘价: %v\n", 
			row["trade_date"], row["open"], row["close"])
	}
}
```

### 使用Bar接口

```go
// 使用通用行情接口
barDf, err := cli.Bar(client.BarParams{
	TsCode:     "000001.SZ",
	StartDate:  "20220101",
	EndDate:    "20220110",
	Freq:       "D",       // D=日线, W=周线, M=月线
	AssetType:  "E",       // E=股票, I=指数, FT=期货, C=数字货币
	AdjustType: "qfq",     // qfq=前复权, hfq=后复权, None=不复权
})
```

### 获取股票基本信息

```go
// 获取股票列表
stockDf, err := cli.GetStockList(client.StockListParams{
	ListStatus: "L",  // L=上市 D=退市 P=暂停上市
	Exchange:   "",   // 交易所 SSE上交所 SZSE深交所 
	IsHS:       "",   // 是否沪深港通标的，N否 H沪股通 S深股通
}, nil)

// 使用简化接口获取所有上市股票
allStocks, err := cli.GetAllListedStocks()
```

### 获取财务数据

```go
// 获取利润表数据
incomeDf, err := cli.GetIncome(client.IncomeParams{
	TSCode:     "000001.SZ",
	Period:     "20211231",
	ReportType: "1", // 合并报表
}, nil)

// 获取资产负债表数据
bsDf, err := cli.GetBalanceSheet(client.BalanceSheetParams{
	TSCode:     "000001.SZ",
	Period:     "20211231",
}, nil)
```

## 已实现接口

### 基础数据

- 股票列表: `stock_basic`
- 交易日历: `trade_cal`
- 股票曾用名: `namechange`
- 沪深股通成份股: `hs_const`
- 上市公司基本信息: `stock_company`
- IPO新股上市: `new_share`

### 行情数据

- 日线行情: `daily`
- 周线行情: `weekly`
- 月线行情: `monthly`
- 分钟线行情: `stk_mins`
- 复权因子: `adj_factor`

### 财务数据

- 利润表: `income`
- 资产负债表: `balancesheet`

## 丰富功能

- 复权处理：支持前复权、后复权
- 均线计算：支持自定义周期的价格和成交量均线
- 通用查询：支持全部TuShare原生接口
- 日志系统：支持多级别日志，可定制输出格式和目的地

## 进度追踪

请查看`PROGRESS.md`文件了解更详细的项目进度。

## 接口文档

请查看`API.md`文件了解详细的接口定义和使用方法。

## 贡献代码

欢迎提交Pull Request或Issue。

## 许可证

MIT
