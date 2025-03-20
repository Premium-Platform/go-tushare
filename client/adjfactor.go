package client

import (
	"github.com/Premium-Platform/go-tushare/pkg/types"
)

// AdjFactorParams 复权因子查询参数
type AdjFactorParams struct {
	TSCode    string `json:"ts_code"`    // 股票代码
	TradeDate string `json:"trade_date"` // 交易日期
	StartDate string `json:"start_date"` // 开始日期
	EndDate   string `json:"end_date"`   // 结束日期
}

// AdjFactorField 复权因子字段常量
var AdjFactorField = struct {
	TSCode    string
	TradeDate string
	AdjFactor string
}{
	TSCode:    "ts_code",    // TS代码
	TradeDate: "trade_date", // 交易日期
	AdjFactor: "adj_factor", // 复权因子
}

// GetAdjFactor 获取复权因子
//
// 接口参数：
// - ts_code: 股票代码
// - trade_date: 交易日期
// - start_date: 开始日期
// - end_date: 结束日期
//
// 返回字段：
// - ts_code: 股票代码
// - trade_date: 交易日期
// - adj_factor: 复权因子
func (c *Client) GetAdjFactor(params AdjFactorParams, fields []string) (*types.DataFrame, error) {
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

	// 调用通用查询接口
	return c.Query("adj_factor", reqParams, fields)
}

// GetStockAdjFactor 获取指定股票的复权因子（简化接口）
func (c *Client) GetStockAdjFactor(tsCode string, startDate string, endDate string) (*types.DataFrame, error) {
	return c.GetAdjFactor(AdjFactorParams{
		TSCode:    tsCode,
		StartDate: startDate,
		EndDate:   endDate,
	}, nil)
}

// GetDayAdjFactor 获取某一天的复权因子（简化接口）
func (c *Client) GetDayAdjFactor(tradeDate string) (*types.DataFrame, error) {
	return c.GetAdjFactor(AdjFactorParams{
		TradeDate: tradeDate,
	}, nil)
}

// CommonAdjFactorFields 返回常用的复权因子字段列表
func (c *Client) CommonAdjFactorFields() []string {
	return []string{
		AdjFactorField.TSCode,
		AdjFactorField.TradeDate,
		AdjFactorField.AdjFactor,
	}
}
