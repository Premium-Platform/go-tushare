package client

import (
	"github.com/Premium-Platform/go-tushare/pkg/types"
)

// NewShareParams IPO新股上市查询参数
type NewShareParams struct {
	StartDate string `json:"start_date"` // 上网发行开始日期
	EndDate   string `json:"end_date"`   // 上网发行结束日期
}

// NewShareField IPO新股上市字段常量
var NewShareField = struct {
	TsCode       string
	SubCode      string
	Name         string
	IPODate      string
	IssueDate    string
	DelistDate   string
	Amount       string
	MarketAmount string
	Price        string
	PE           string
	LimitAmount  string
	Funds        string
	Ballot       string
}{
	TsCode:       "ts_code",       // TS股票代码
	SubCode:      "sub_code",      // 申购代码
	Name:         "name",          // 名称
	IPODate:      "ipo_date",      // 上网发行日期
	IssueDate:    "issue_date",    // 上市日期
	DelistDate:   "delist_date",   // 退市日期
	Amount:       "amount",        // 发行总量(万股)
	MarketAmount: "market_amount", // 上网发行总量(万股)
	Price:        "price",         // 发行价格
	PE:           "pe",            // 市盈率
	LimitAmount:  "limit_amount",  // 个人申购上限(万股)
	Funds:        "funds",         // 募集资金(亿元)
	Ballot:       "ballot",        // 中签率
}

// GetNewShare 获取IPO新股上市数据
//
// 接口参数：
// - start_date: 上网发行开始日期
// - end_date: 上网发行结束日期
//
// 返回字段：
// - ts_code: TS股票代码
// - sub_code: 申购代码
// - name: 名称
// - ipo_date: 上网发行日期
// - issue_date: 上市日期
// - delist_date: 退市日期
// - amount: 发行总量(万股)
// - market_amount: 上网发行总量(万股)
// - price: 发行价格
// - pe: 市盈率
// - limit_amount: 个人申购上限(万股)
// - funds: 募集资金(亿元)
// - ballot: 中签率
func (c *Client) GetNewShare(params NewShareParams, fields []string) (*types.DataFrame, error) {
	// 构建请求参数
	reqParams := map[string]interface{}{}

	if params.StartDate != "" {
		reqParams["start_date"] = params.StartDate
	}

	if params.EndDate != "" {
		reqParams["end_date"] = params.EndDate
	}

	// 调用通用查询接口
	return c.Query("new_share", reqParams, fields)
}

// GetRecentNewShares 获取近期新股上市（简化接口）
func (c *Client) GetRecentNewShares() (*types.DataFrame, error) {
	return c.GetNewShare(NewShareParams{}, []string{})
}

// GetNewSharesByPeriod 获取指定时间段内的新股上市信息
func (c *Client) GetNewSharesByPeriod(startDate, endDate string) (*types.DataFrame, error) {
	return c.GetNewShare(NewShareParams{
		StartDate: startDate,
		EndDate:   endDate,
	}, []string{})
}

// CommonNewShareFields 返回IPO新股上市常用字段
func (c *Client) CommonNewShareFields() []string {
	return []string{
		NewShareField.TsCode,
		NewShareField.SubCode,
		NewShareField.Name,
		NewShareField.IPODate,
		NewShareField.IssueDate,
		NewShareField.Price,
		NewShareField.PE,
		NewShareField.Funds,
		NewShareField.Ballot,
	}
}
