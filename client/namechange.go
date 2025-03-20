package client

import (
	"github.com/Premium-Platform/go-tushare/pkg/types"
)

// NameChangeParams 股票曾用名查询参数
type NameChangeParams struct {
	TsCode    string `json:"ts_code"`    // TS代码
	StartDate string `json:"start_date"` // 公告开始日期 (YYYYMMDD)
	EndDate   string `json:"end_date"`   // 公告结束日期 (YYYYMMDD)
}

// NameChangeField 股票曾用名字段常量
var NameChangeField = struct {
	TsCode       string
	Name         string
	StartDate    string
	EndDate      string
	AnnDate      string
	ChangeReason string
}{
	TsCode:       "ts_code",       // TS代码
	Name:         "name",          // 证券名称
	StartDate:    "start_date",    // 开始日期
	EndDate:      "end_date",      // 结束日期
	AnnDate:      "ann_date",      // 公告日期
	ChangeReason: "change_reason", // 变更原因
}

// GetNameChange 获取股票曾用名
//
// 接口参数：
// - ts_code: TS代码
// - start_date: 公告开始日期 (YYYYMMDD)
// - end_date: 公告结束日期 (YYYYMMDD)
//
// 返回字段：
// - ts_code: TS代码
// - name: 证券名称
// - start_date: 开始日期
// - end_date: 结束日期
// - ann_date: 公告日期
// - change_reason: 变更原因
func (c *Client) GetNameChange(params NameChangeParams, fields []string) (*types.DataFrame, error) {
	// 构建请求参数
	reqParams := map[string]interface{}{}

	if params.TsCode != "" {
		reqParams["ts_code"] = params.TsCode
	}

	if params.StartDate != "" {
		reqParams["start_date"] = params.StartDate
	}

	if params.EndDate != "" {
		reqParams["end_date"] = params.EndDate
	}

	// 调用通用查询接口
	return c.Query("namechange", reqParams, fields)
}

// GetStockNameHistory 获取单个股票的名称变更历史（简化接口）
func (c *Client) GetStockNameHistory(tsCode string) (*types.DataFrame, error) {
	return c.GetNameChange(NameChangeParams{
		TsCode: tsCode,
	}, []string{})
}

// GetNameChangeInPeriod 获取指定时间段内的股票名称变更记录
func (c *Client) GetNameChangeInPeriod(startDate, endDate string) (*types.DataFrame, error) {
	return c.GetNameChange(NameChangeParams{
		StartDate: startDate,
		EndDate:   endDate,
	}, []string{})
}

// CommonNameChangeFields 返回股票曾用名常用字段
func (c *Client) CommonNameChangeFields() []string {
	return []string{
		NameChangeField.TsCode,
		NameChangeField.Name,
		NameChangeField.StartDate,
		NameChangeField.EndDate,
		NameChangeField.AnnDate,
		NameChangeField.ChangeReason,
	}
}
