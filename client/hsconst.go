package client

import (
	"github.com/Premium-Platform/go-tushare/pkg/types"
)

// HSConstParams 沪深股通成份股查询参数
type HSConstParams struct {
	HSType string `json:"hs_type"` // 沪深港通类型，SH沪股通，SZ深股通
	IsNew  string `json:"is_new"`  // 是否最新，1是 0否 (默认1)
}

// HSConstField 沪深股通成份股字段常量
var HSConstField = struct {
	TsCode  string
	HSType  string
	InDate  string
	OutDate string
	IsNew   string
}{
	TsCode:  "ts_code",  // TS代码
	HSType:  "hs_type",  // 沪深港通类型，SH沪股通，SZ深股通
	InDate:  "in_date",  // 纳入日期
	OutDate: "out_date", // 剔除日期
	IsNew:   "is_new",   // 是否最新，1是 0否
}

// GetHSConst 获取沪深股通成份股
//
// 接口参数：
// - hs_type: 沪深港通类型，SH沪股通，SZ深股通
// - is_new: 是否最新，1是 0否 (默认1)
//
// 返回字段：
// - ts_code: TS代码
// - hs_type: 沪深港通类型，SH沪股通，SZ深股通
// - in_date: 纳入日期
// - out_date: 剔除日期
// - is_new: 是否最新，1是 0否
func (c *Client) GetHSConst(params HSConstParams, fields []string) (*types.DataFrame, error) {
	// 构建请求参数
	reqParams := map[string]interface{}{}

	if params.HSType != "" {
		reqParams["hs_type"] = params.HSType
	}

	if params.IsNew != "" {
		reqParams["is_new"] = params.IsNew
	} else {
		// 默认获取最新
		reqParams["is_new"] = "1"
	}

	// 调用通用查询接口
	return c.Query("hs_const", reqParams, fields)
}

// GetSHConst 获取沪股通成份股（简化接口）
func (c *Client) GetSHConst() (*types.DataFrame, error) {
	return c.GetHSConst(HSConstParams{
		HSType: "SH",
		IsNew:  "1",
	}, []string{})
}

// GetSZConst 获取深股通成份股（简化接口）
func (c *Client) GetSZConst() (*types.DataFrame, error) {
	return c.GetHSConst(HSConstParams{
		HSType: "SZ",
		IsNew:  "1",
	}, []string{})
}

// GetHSConstHistory 获取沪深港通成份股历史记录
func (c *Client) GetHSConstHistory(hsType string) (*types.DataFrame, error) {
	return c.GetHSConst(HSConstParams{
		HSType: hsType,
		IsNew:  "0",
	}, []string{})
}

// CommonHSConstFields 返回沪深股通成份股常用字段
func (c *Client) CommonHSConstFields() []string {
	return []string{
		HSConstField.TsCode,
		HSConstField.HSType,
		HSConstField.InDate,
		HSConstField.OutDate,
		HSConstField.IsNew,
	}
}
