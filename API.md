# Go-TuShare API文档

本文档说明Go-TuShare SDK的API接口定义。

## 核心接口

### 初始化客户端

```go
// 创建新客户端
client.New(token string) *Client

// 设置超时时间（可选）
client.SetTimeout(timeout time.Duration)

// 设置令牌
client.SetToken(token string)

// 获取令牌
client.GetToken() string
```

### 通用查询接口

```go
// 通用API查询
client.Query(apiName string, params map[string]interface{}, fields []string) (*DataFrame, error)

// 数据使用示例
df, err := client.Query("daily", map[string]interface{}{
    "ts_code": "000001.SZ",
    "start_date": "20180101",
    "end_date": "20180115",
}, []string{})

if err != nil {
    // 处理错误
}

// 访问数据
for _, row := range df.Rows {
    // 处理每一行数据
}
```

### 行情数据通用接口

```go
// 行情数据通用接口 (类似pro_bar)
client.Bar(params BarParams) (*DataFrame, error)

// BarParams结构体
type BarParams struct {
    TsCode       string   // 证券代码
    StartDate    string   // 开始日期 YYYYMMDD
    EndDate      string   // 结束日期 YYYYMMDD
    Freq         string   // 周期：D=日线 W=周线 M=月线
    AssetType    string   // 资产类别：E=股票 I=指数 C=数字货币 FT=期货
    Exchange     string   // 交易所
    AdjustType   string   // 复权类型：None=不复权 qfq=前复权 hfq=后复权
    MA           []int    // 均线
    Factors      []string // 因子数据
    ContractType string   // 合约类型
}
```

## 数据模型

### DataFrame

数据返回的基础模型

```go
// DataFrame 结构体
type DataFrame struct {
    Columns []string                   // 列名
    Rows    []map[string]interface{}   // 行数据
}

// 获取指定列数据
func (df *DataFrame) GetColumn(name string) []interface{}

// 转换为JSON
func (df *DataFrame) ToJSON() ([]byte, error)

// 转换为CSV
func (df *DataFrame) ToCSV() ([]byte, error)
```

## 接口列表

### 基础数据

- 股票列表: `stock_basic`
- 交易日历: `trade_cal`
- 股票曾用名: `namechange`
- 沪深股通成份股: `hs_const`

### 行情数据

- 日线行情: `daily`
- 周线行情: `weekly`
- 月线行情: `monthly`
- 分钟线行情: `stk_mins`
- 复权因子: `adj_factor`

### 财务数据

- 利润表: `income`
- 资产负债表: `balancesheet`
- 现金流量表: `cashflow` (未实现)
- 业绩预告: `forecast` (未实现)
- 业绩快报: `express` (未实现)

### 指数数据

- 指数日线行情: `index_daily` (未实现)
- 指数周线行情: `index_weekly` (未实现)
- 指数月线行情: `index_monthly` (未实现)
- 指数成分和权重: `index_weight` (未实现)
- 大盘指数每日指标: `index_dailybasic` (未实现)

## 新增接口示例

### 复权因子接口

```go
// 获取复权因子
adjDf, err := cli.GetAdjFactor(client.AdjFactorParams{
    TSCode:     "000001.SZ",
    StartDate:  "20220101",
    EndDate:    "20220110",
}, nil)

// 使用简化接口获取特定股票的复权因子
adjDf, err := cli.GetStockAdjFactor("000001.SZ", "20220101", "20220110")
```

### 分钟线数据接口

```go
// 获取分钟线数据
minDf, err := cli.GetStockMinute(client.MinuteParams{
    TSCode:     "000001.SZ",
    TradeDate:  "20220110",
    Freq:       "5", // 5分钟线
}, nil)

// 使用简化接口获取5分钟线数据
min5Df, err := cli.Get5MinLine("000001.SZ", "20220110")
```

### 财务数据接口

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

// 使用简化接口获取最新季度财务数据
latestIncome, err := cli.GetLatestIncome("000001.SZ")
latestBS, err := cli.GetLatestBalanceSheet("000001.SZ")
```

更多接口详情将在后续开发中添加。 