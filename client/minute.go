package client

import (
	"github.com/Premium-Platform/go-tushare/pkg/types"
)

// MinuteParams 分钟线行情查询参数
type MinuteParams struct {
	TSCode    string `json:"ts_code"`    // 股票代码
	TradeDate string `json:"trade_date"` // 交易日期
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
	StartTime string `json:"start_time"` // 开始时间
	EndTime   string `json:"end_time"`   // 结束时间
	Freq      string `json:"freq"`       // 频率，1，5，15，30，60分钟
}

// MinuteField 分钟线行情字段常量
var MinuteField = struct {
	TSCode      string
	TradeTime   string
	TradeDate   string
	Open        string
	High        string
	Low         string
	Close       string
	PreClose    string
	Vol         string
	Amount      string
	MA5         string
	MA10        string
	MA20        string
	VolumeRatio string
}{
	TSCode:      "ts_code",      // TS代码
	TradeTime:   "trade_time",   // 交易时间
	TradeDate:   "trade_date",   // 交易日期
	Open:        "open",         // 开盘价
	High:        "high",         // 最高价
	Low:         "low",          // 最低价
	Close:       "close",        // 收盘价
	PreClose:    "pre_close",    // 昨收价
	Vol:         "vol",          // 成交量
	Amount:      "amount",       // 成交额
	MA5:         "ma5",          // 5分钟均线
	MA10:        "ma10",         // 10分钟均线
	MA20:        "ma20",         // 20分钟均线
	VolumeRatio: "volume_ratio", // 量比
}

// GetStockMinute 获取分钟线行情数据
//
// 接口参数：
// - ts_code: 股票代码
// - trade_date: 交易日期
// - start_date: 开始日期
// - end_date: 结束日期
// - start_time: 开始时间
// - end_time: 结束时间
// - freq: 频率，1，5，15，30，60分钟
//
// 返回字段：
// - ts_code: 股票代码
// - trade_time: 交易时间
// - open: 开盘价
// - high: 最高价
// - low: 最低价
// - close: 收盘价
// - pre_close: 昨收价
// - vol: 成交量
// - amount: 成交额
func (c *Client) GetStockMinute(params MinuteParams, fields []string) (*types.DataFrame, error) {
	// 构建请求参数
	reqParams := map[string]interface{}{}

	if params.TSCode != "" {
		reqParams["ts_code"] = params.TSCode
	}

	if params.TradeDate != "" {
		reqParams["trade_date"] = params.TradeDate
	}

	if params.StartDate != "" {
		reqParams["start_date"] = params.StartDate
	}

	if params.EndDate != "" {
		reqParams["end_date"] = params.EndDate
	}

	if params.StartTime != "" {
		reqParams["start_time"] = params.StartTime
	}

	if params.EndTime != "" {
		reqParams["end_time"] = params.EndTime
	}

	// 默认为1分钟
	if params.Freq == "" {
		params.Freq = "1"
	}
	reqParams["freq"] = params.Freq

	// 调用通用查询接口
	return c.Query("stk_mins", reqParams, fields)
}

// Get1MinLine 获取1分钟线数据（简化接口）
func (c *Client) Get1MinLine(tsCode string, tradeDate string) (*types.DataFrame, error) {
	return c.GetStockMinute(MinuteParams{
		TSCode:    tsCode,
		TradeDate: tradeDate,
		Freq:      "1",
	}, nil)
}

// Get5MinLine 获取5分钟线数据（简化接口）
func (c *Client) Get5MinLine(tsCode string, tradeDate string) (*types.DataFrame, error) {
	return c.GetStockMinute(MinuteParams{
		TSCode:    tsCode,
		TradeDate: tradeDate,
		Freq:      "5",
	}, nil)
}

// Get15MinLine 获取15分钟线数据（简化接口）
func (c *Client) Get15MinLine(tsCode string, tradeDate string) (*types.DataFrame, error) {
	return c.GetStockMinute(MinuteParams{
		TSCode:    tsCode,
		TradeDate: tradeDate,
		Freq:      "15",
	}, nil)
}

// Get30MinLine 获取30分钟线数据（简化接口）
func (c *Client) Get30MinLine(tsCode string, tradeDate string) (*types.DataFrame, error) {
	return c.GetStockMinute(MinuteParams{
		TSCode:    tsCode,
		TradeDate: tradeDate,
		Freq:      "30",
	}, nil)
}

// Get60MinLine 获取60分钟线数据（简化接口）
func (c *Client) Get60MinLine(tsCode string, tradeDate string) (*types.DataFrame, error) {
	return c.GetStockMinute(MinuteParams{
		TSCode:    tsCode,
		TradeDate: tradeDate,
		Freq:      "60",
	}, nil)
}

// CommonMinuteFields 返回常用的分钟线字段列表
func (c *Client) CommonMinuteFields() []string {
	return []string{
		MinuteField.TSCode,
		MinuteField.TradeTime,
		MinuteField.Open,
		MinuteField.High,
		MinuteField.Low,
		MinuteField.Close,
		MinuteField.Vol,
		MinuteField.Amount,
	}
}
