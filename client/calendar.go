package client

import (
	"github.com/Premium-Platform/go-tushare/pkg/types"
)

// TradeCalParams 交易日历查询参数
type TradeCalParams struct {
	Exchange  string `json:"exchange"`   // 交易所 SSE上交所,SZSE深交所,CFFEX中金所,SHFE上期所,CZCE郑商所,DCE大商所,INE上能源
	StartDate string `json:"start_date"` // 开始日期 (格式：YYYYMMDD)
	EndDate   string `json:"end_date"`   // 结束日期 (格式：YYYYMMDD)
	IsOpen    string `json:"is_open"`    // 是否交易 '0'休市 '1'交易
}

// TradeCalField 交易日历字段常量
var TradeCalField = struct {
	Exchange     string
	CalDate      string
	IsOpen       string
	PretradeDate string
}{
	Exchange:     "exchange",      // 交易所 SSE上交所 SZSE深交所
	CalDate:      "cal_date",      // 日历日期
	IsOpen:       "is_open",       // 是否交易 0休市 1交易
	PretradeDate: "pretrade_date", // 上一个交易日
}

// GetTradeCal 获取交易日历
//
// 接口参数：
// - exchange: 交易所 SSE上交所,SZSE深交所,CFFEX中金所,SHFE上期所,CZCE郑商所,DCE大商所,INE上能源
// - start_date: 开始日期 (格式：YYYYMMDD)
// - end_date: 结束日期 (格式：YYYYMMDD)
// - is_open: 是否交易 '0'休市 '1'交易
//
// 返回字段：
// - exchange: 交易所 SSE上交所 SZSE深交所
// - cal_date: 日历日期
// - is_open: 是否交易 0休市 1交易
// - pretrade_date: 上一个交易日
func (c *Client) GetTradeCal(params TradeCalParams, fields []string) (*types.DataFrame, error) {
	// 构建请求参数
	reqParams := map[string]interface{}{}

	if params.Exchange != "" {
		reqParams["exchange"] = params.Exchange
	} else {
		// 默认上交所
		reqParams["exchange"] = "SSE"
	}

	if params.StartDate != "" {
		reqParams["start_date"] = params.StartDate
	}

	if params.EndDate != "" {
		reqParams["end_date"] = params.EndDate
	}

	if params.IsOpen != "" {
		reqParams["is_open"] = params.IsOpen
	}

	// 调用通用查询接口
	return c.Query("trade_cal", reqParams, fields)
}

// GetTradeCalWithDefault 获取指定范围的交易日历（简化接口）
func (c *Client) GetTradeCalWithDefault(startDate, endDate string) (*types.DataFrame, error) {
	return c.GetTradeCal(TradeCalParams{
		Exchange:  "SSE",
		StartDate: startDate,
		EndDate:   endDate,
	}, []string{})
}

// GetTradeDays 获取交易日历中的交易日
func (c *Client) GetTradeDays(startDate, endDate string) (*types.DataFrame, error) {
	return c.GetTradeCal(TradeCalParams{
		Exchange:  "SSE",
		StartDate: startDate,
		EndDate:   endDate,
		IsOpen:    "1",
	}, []string{})
}

// GetSSETradeCal 获取上交所交易日历
func (c *Client) GetSSETradeCal(startDate, endDate string) (*types.DataFrame, error) {
	return c.GetTradeCal(TradeCalParams{
		Exchange:  "SSE",
		StartDate: startDate,
		EndDate:   endDate,
	}, []string{})
}

// GetSZSETradeCal 获取深交所交易日历
func (c *Client) GetSZSETradeCal(startDate, endDate string) (*types.DataFrame, error) {
	return c.GetTradeCal(TradeCalParams{
		Exchange:  "SZSE",
		StartDate: startDate,
		EndDate:   endDate,
	}, []string{})
}

// CommonTradeCalFields 返回交易日历常用字段
func (c *Client) CommonTradeCalFields() []string {
	return []string{
		TradeCalField.Exchange,
		TradeCalField.CalDate,
		TradeCalField.IsOpen,
		TradeCalField.PretradeDate,
	}
}
