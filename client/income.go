package client

import (
	"github.com/Premium-Platform/go-tushare/pkg/types"
)

// IncomeParams 利润表查询参数
type IncomeParams struct {
	TSCode     string `json:"ts_code"`     // 股票代码
	AnnDate    string `json:"ann_date"`    // 公告日期
	StartDate  string `json:"start_date"`  // 公告开始日期
	EndDate    string `json:"end_date"`    // 公告结束日期
	Period     string `json:"period"`      // 报告期
	ReportType string `json:"report_type"` // 报告类型
	CompType   string `json:"comp_type"`   // 公司类型
}

// IncomeField 利润表字段常量
var IncomeField = struct {
	TSCode            string
	AnnDate           string
	FAnnDate          string
	EndDate           string
	ReportType        string
	CompType          string
	BasicEPS          string
	DilutedEPS        string
	TotalRevenue      string
	Revenue           string
	IntIncome         string
	PremEarned        string
	CommIncome        string
	NCommIncome       string
	NSecTradingIncome string
	NSecInvestIncome  string
	NIntNIntIncome    string
	InvestIncome      string
	AssInvestIncome   string
	ForexGain         string
	OperateProfit     string
	NonOperIncome     string
	NonOperExp        string
	NonCurAssetEarned string
	TotalProfit       string
	IncomeTax         string
	NIncome           string
	NIncomeAttrP      string
	MinorityGain      string
	OthComprIncome    string
	TComprIncome      string
	ComprIncAttrP     string
	ComprIncAttrMS    string
	EBIT              string
	EBITDA            string
	InsurExp          string
	UndistProfit      string
	DistableProfit    string
}{
	TSCode:            "ts_code",              // TS代码
	AnnDate:           "ann_date",             // 公告日期
	FAnnDate:          "f_ann_date",           // 实际公告日期
	EndDate:           "end_date",             // 报告期
	ReportType:        "report_type",          // 报告类型
	CompType:          "comp_type",            // 公司类型
	BasicEPS:          "basic_eps",            // 基本每股收益
	DilutedEPS:        "diluted_eps",          // 稀释每股收益
	TotalRevenue:      "total_revenue",        // 营业总收入
	Revenue:           "revenue",              // 营业收入
	IntIncome:         "int_income",           // 利息收入
	PremEarned:        "prem_earned",          // 已赚保费
	CommIncome:        "comm_income",          // 手续费及佣金收入
	NCommIncome:       "n_comm_income",        // 手续费及佣金净收入
	NSecTradingIncome: "n_sec_trading_income", // 证券买卖业务净收入
	NSecInvestIncome:  "n_sec_invest_income",  // 证券投资收益
	NIntNIntIncome:    "n_int_n_int_income",   // 利息净收入
	InvestIncome:      "invest_income",        // 投资收益
	AssInvestIncome:   "ass_invest_income",    // 对联营企业和合营企业的投资收益
	ForexGain:         "forex_gain",           // 汇兑收益
	OperateProfit:     "operate_profit",       // 营业利润
	NonOperIncome:     "non_oper_income",      // 营业外收入
	NonOperExp:        "non_oper_exp",         // 营业外支出
	NonCurAssetEarned: "non_cur_asset_earned", // 非流动资产处置利得
	TotalProfit:       "total_profit",         // 利润总额
	IncomeTax:         "income_tax",           // 所得税费用
	NIncome:           "n_income",             // 净利润
	NIncomeAttrP:      "n_income_attr_p",      // 归属于母公司股东的净利润
	MinorityGain:      "minority_gain",        // 少数股东损益
	OthComprIncome:    "oth_compr_income",     // 其他综合收益
	TComprIncome:      "t_compr_income",       // 综合收益总额
	ComprIncAttrP:     "compr_inc_attr_p",     // 归属于母公司所有者的综合收益总额
	ComprIncAttrMS:    "compr_inc_attr_ms",    // 归属于少数股东的综合收益总额
	EBIT:              "ebit",                 // 息税前利润
	EBITDA:            "ebitda",               // 息税折旧摊销前利润
	InsurExp:          "insur_exp",            // 保险业务支出
	UndistProfit:      "undist_profit",        // 年初未分配利润
	DistableProfit:    "distable_profit",      // 可分配利润
}

// GetIncome 获取利润表数据
//
// 接口参数：
// - ts_code: 股票代码
// - ann_date: 公告日期
// - start_date: 公告开始日期
// - end_date: 公告结束日期
// - period: 报告期
// - report_type: 报告类型  1合并报表 2单季合并 3调整单季合并表 4调整合并报表 5调整前合并报表 6母公司报表 7母公司单季表 8 母公司调整单季表 9母公司调整表 10母公司调整前报表 11调整前合并报表 12母公司调整前报表
// - comp_type: 公司类型  1一般工商业 2银行 3保险 4证券
//
// 返回字段：
// - ts_code: 股票代码
// - ann_date: 公告日期
// - f_ann_date: 实际公告日期
// - end_date: 报告期
// - report_type: 报告类型
// - comp_type: 公司类型
// - basic_eps: 基本每股收益
// - diluted_eps: 稀释每股收益
// - total_revenue: 营业总收入
// - revenue: 营业收入
// - int_income: 利息收入
// - prem_earned: 已赚保费
// - comm_income: 手续费及佣金收入
// - n_comm_income: 手续费及佣金净收入
// - n_sec_trading_income: 证券买卖业务净收入
// - n_sec_invest_income: 证券投资收益
// - n_int_n_int_income: 利息净收入
// - invest_income: 投资收益
// - ass_invest_income: 对联营企业和合营企业的投资收益
// - forex_gain: 汇兑收益
// - operate_profit: 营业利润
// - non_oper_income: 营业外收入
// - non_oper_exp: 营业外支出
// - non_cur_asset_earned: 非流动资产处置利得
// - total_profit: 利润总额
// - income_tax: 所得税费用
// - n_income: 净利润
// - n_income_attr_p: 归属于母公司股东的净利润
// - minority_gain: 少数股东损益
// - oth_compr_income: 其他综合收益
// - t_compr_income: 综合收益总额
// - compr_inc_attr_p: 归属于母公司所有者的综合收益总额
// - compr_inc_attr_ms: 归属于少数股东的综合收益总额
// - ebit: 息税前利润
// - ebitda: 息税折旧摊销前利润
// - insur_exp: 保险业务支出
// - undist_profit: 年初未分配利润
// - distable_profit: 可分配利润
func (c *Client) GetIncome(params IncomeParams, fields []string) (*types.DataFrame, error) {
	// 构建请求参数
	reqParams := map[string]interface{}{}

	if params.TSCode != "" {
		reqParams["ts_code"] = params.TSCode
	}

	if params.AnnDate != "" {
		reqParams["ann_date"] = params.AnnDate
	}

	if params.StartDate != "" {
		reqParams["start_date"] = params.StartDate
	}

	if params.EndDate != "" {
		reqParams["end_date"] = params.EndDate
	}

	if params.Period != "" {
		reqParams["period"] = params.Period
	}

	if params.ReportType != "" {
		reqParams["report_type"] = params.ReportType
	}

	if params.CompType != "" {
		reqParams["comp_type"] = params.CompType
	}

	// 调用通用查询接口
	return c.Query("income", reqParams, fields)
}

// GetLatestIncome 获取最新的利润表数据（简化接口）
func (c *Client) GetLatestIncome(tsCode string) (*types.DataFrame, error) {
	return c.GetIncome(IncomeParams{
		TSCode:     tsCode,
		ReportType: "1", // 默认获取合并报表
	}, nil)
}

// GetYearIncome 获取年度利润表数据（简化接口）
func (c *Client) GetYearIncome(tsCode string, year string) (*types.DataFrame, error) {
	return c.GetIncome(IncomeParams{
		TSCode:     tsCode,
		Period:     year + "1231", // 年度报表日期
		ReportType: "1",           // 合并报表
	}, nil)
}

// GetQuarterIncome 获取季度利润表数据（简化接口）
func (c *Client) GetQuarterIncome(tsCode string, yearQuarter string) (*types.DataFrame, error) {
	return c.GetIncome(IncomeParams{
		TSCode:     tsCode,
		Period:     yearQuarter, // 如"20211231", "20220331"
		ReportType: "1",         // 合并报表
	}, nil)
}

// CommonIncomeFields 返回常用的利润表字段列表
func (c *Client) CommonIncomeFields() []string {
	return []string{
		IncomeField.TSCode,
		IncomeField.EndDate,
		IncomeField.AnnDate,
		IncomeField.TotalRevenue,
		IncomeField.OperateProfit,
		IncomeField.TotalProfit,
		IncomeField.NIncome,
		IncomeField.NIncomeAttrP,
		IncomeField.BasicEPS,
		IncomeField.DilutedEPS,
	}
}
