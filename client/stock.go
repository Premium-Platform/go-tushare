package client

import (
	"github.com/Premium-Platform/go-tushare/pkg/types"
)

// StockBasicParams 股票列表查询参数
type StockBasicParams struct {
	IsHS       string `json:"is_hs"`       // 是否沪深港通标的，N否 H沪股通 S深股通
	ListStatus string `json:"list_status"` // 上市状态：L上市 D退市 P暂停上市，默认是L
	Exchange   string `json:"exchange"`    // 交易所 SSE上交所 SZSE深交所 BSE北交所
	Market     string `json:"market"`      // 市场类别 （主板/创业板/科创板/CDR/北交所）
}

// StockBasicField 股票列表字段常量
var StockBasicField = struct {
	TSCode     string
	Symbol     string
	Name       string
	Area       string
	Industry   string
	Fullname   string
	Enname     string
	Cnspell    string
	Market     string
	Exchange   string
	CurrType   string
	ListStatus string
	ListDate   string
	DelistDate string
	IsHS       string
	ActName    string
	ActEntType string
}{
	TSCode:     "ts_code",      // TS代码
	Symbol:     "symbol",       // 股票代码
	Name:       "name",         // 股票名称
	Area:       "area",         // 地域
	Industry:   "industry",     // 所属行业
	Fullname:   "fullname",     // 股票全称
	Enname:     "enname",       // 英文全称
	Cnspell:    "cnspell",      // 拼音缩写
	Market:     "market",       // 市场类型（主板/创业板/科创板/CDR）
	Exchange:   "exchange",     // 交易所代码
	CurrType:   "curr_type",    // 交易货币
	ListStatus: "list_status",  // 上市状态
	ListDate:   "list_date",    // 上市日期
	DelistDate: "delist_date",  // 退市日期
	IsHS:       "is_hs",        // 是否沪深港通标的
	ActName:    "act_name",     // 实控人名称
	ActEntType: "act_ent_type", // 实控人企业性质
}

// GetStockBasic 获取股票列表
//
// 接口参数：
// - exchange: 交易所代码，SSE上交所 SZSE深交所 BSE北交所
// - list_status: 上市状态 L上市 D退市 P暂停上市，默认是L
// - is_hs: 是否沪深港通标的，N否 H沪股通 S深股通
// - market: 市场类别 （主板/创业板/科创板/CDR/北交所）
//
// 返回字段：
// - ts_code: TS代码
// - symbol: 股票代码
// - name: 股票名称
// - area: 地域
// - industry: 所属行业
// - fullname: 股票全称
// - enname: 英文全称
// - cnspell: 拼音缩写
// - market: 市场类型
// - exchange: 交易所代码
// - curr_type: 交易货币
// - list_status: 上市状态
// - list_date: 上市日期
// - delist_date: 退市日期
// - is_hs: 是否沪深港通标的
// - act_name: 实控人名称
// - act_ent_type: 实控人企业性质
func (c *Client) GetStockBasic(params StockBasicParams, fields []string) (*types.DataFrame, error) {
	// 构建请求参数
	reqParams := map[string]interface{}{}

	if params.IsHS != "" {
		reqParams["is_hs"] = params.IsHS
	}

	if params.ListStatus != "" {
		reqParams["list_status"] = params.ListStatus
	} else {
		// 默认获取上市股票
		reqParams["list_status"] = "L"
	}

	if params.Exchange != "" {
		reqParams["exchange"] = params.Exchange
	}

	if params.Market != "" {
		reqParams["market"] = params.Market
	}

	// 调用通用查询接口
	return c.Query("stock_basic", reqParams, fields)
}

// ListStocks 获取所有上市股票列表（简化接口）
func (c *Client) ListStocks() (*types.DataFrame, error) {
	return c.GetStockBasic(StockBasicParams{ListStatus: "L"}, []string{})
}

// CommonFields 返回常用的字段列表
func (c *Client) CommonStockFields() []string {
	return []string{
		StockBasicField.TSCode,
		StockBasicField.Symbol,
		StockBasicField.Name,
		StockBasicField.Area,
		StockBasicField.Industry,
		StockBasicField.Market,
		StockBasicField.ListDate,
	}
}

// GetMainboardStocks 获取主板股票
func (c *Client) GetMainboardStocks() (*types.DataFrame, error) {
	return c.GetStockBasic(StockBasicParams{
		ListStatus: "L",
		Market:     "主板",
	}, c.CommonStockFields())
}

// GetGEMStocks 获取创业板股票
func (c *Client) GetGEMStocks() (*types.DataFrame, error) {
	return c.GetStockBasic(StockBasicParams{
		ListStatus: "L",
		Market:     "创业板",
	}, c.CommonStockFields())
}

// GetSTARStocks 获取科创板股票
func (c *Client) GetSTARStocks() (*types.DataFrame, error) {
	return c.GetStockBasic(StockBasicParams{
		ListStatus: "L",
		Market:     "科创板",
	}, c.CommonStockFields())
}

// 股票基本信息字段说明
// ts_code        股票代码，如：000001.SZ, 600000.SH
// symbol         股票代码，如：000001，600000
// name           股票名称
// area           地区
// industry       所属行业
// fullname       股票全称
// enname         英文全称
// market         市场类型 （主板/中小板/创业板/科创板）
// exchange       交易所代码，如：SSE上交所 SZSE深交所 HKEX港交所
// curr_type      交易货币，如：CNY人民币 HKD港币 USD美元
// list_status    上市状态：L上市 D退市 P暂停上市
// list_date      上市日期
// delist_date    退市日期
// is_hs          是否沪深港通标的，N否 H沪股通 S深股通
