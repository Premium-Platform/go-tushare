package client

import (
	"github.com/Premium-Platform/go-tushare/pkg/types"
)

// BalanceSheetParams 资产负债表查询参数
type BalanceSheetParams struct {
	TSCode     string `json:"ts_code"`     // 股票代码
	AnnDate    string `json:"ann_date"`    // 公告日期
	StartDate  string `json:"start_date"`  // 公告开始日期
	EndDate    string `json:"end_date"`    // 公告结束日期
	Period     string `json:"period"`      // 报告期
	ReportType string `json:"report_type"` // 报告类型
	CompType   string `json:"comp_type"`   // 公司类型
}

// BalanceSheetField 资产负债表字段常量
var BalanceSheetField = struct {
	TSCode             string
	AnnDate            string
	FAnnDate           string
	EndDate            string
	ReportType         string
	CompType           string
	TotalAssets        string
	TotalLiab          string
	TotalSHEquity      string
	TotalEquityExcMin  string
	TotalCurrAssets    string
	TotalNonCurrAssets string
	TotalCurrLiab      string
	TotalNonCurrLiab   string
	MoneyCapital       string
	TradeAssets        string
	NotesReceiv        string
	AccountsReceiv     string
	OtherReceiv        string
	Inventories        string
	PrepayMents        string
	DividReceiv        string
	LtEquityInvest     string
	TimeDeposits       string
	OtherAssets        string
	FixedAssets        string
	ConstructionMater  string
	ConstructionProg   string
	IntangibleAssets   string
	RightUseAssets     string
	DeferredTaxAssets  string
	CPledgeLoans       string
	ShortBorrows       string
	NotesPayable       string
	AccountsPayable    string
	AdvanceReceipts    string
	ContractLiab       string
	EmployeeBenefits   string
	TaxesPayable       string
	TotalOtherPayable  string
	EstimateLiab       string
	LtBorrow           string
	BondsPayable       string
	LtPayable          string
	SpecialReserve     string
	CapitalRece        string
	SurplusReserve     string
	RetainedEarnings   string
	TreasuryShare      string
}{
	TSCode:             "ts_code",                    // TS代码
	AnnDate:            "ann_date",                   // 公告日期
	FAnnDate:           "f_ann_date",                 // 实际公告日期
	EndDate:            "end_date",                   // 报告期
	ReportType:         "report_type",                // 报告类型
	CompType:           "comp_type",                  // 公司类型
	TotalAssets:        "total_assets",               // 总资产
	TotalLiab:          "total_liab",                 // 总负债
	TotalSHEquity:      "total_hldr_eqy_inc_min_int", // 股东权益合计(含少数股东权益)
	TotalEquityExcMin:  "total_hldr_eqy_exc_min_int", // 股东权益合计(不含少数股东权益)
	TotalCurrAssets:    "total_cur_assets",           // 流动资产合计
	TotalNonCurrAssets: "total_non_cur_assets",       // 非流动资产合计
	TotalCurrLiab:      "total_cur_liab",             // 流动负债合计
	TotalNonCurrLiab:   "total_non_cur_liab",         // 非流动负债合计
	MoneyCapital:       "money_cap",                  // 货币资金
	TradeAssets:        "trad_asset",                 // 交易性金融资产
	NotesReceiv:        "notes_receiv",               // 应收票据
	AccountsReceiv:     "accounts_receiv",            // 应收账款
	OtherReceiv:        "oth_receiv",                 // 其他应收款
	Inventories:        "inventories",                // 存货
	PrepayMents:        "prepayment",                 // 预付款项
	DividReceiv:        "div_receiv",                 // 应收股利
	LtEquityInvest:     "lt_eqt_invest",              // 长期股权投资
	TimeDeposits:       "time_deposits",              // 定期存款
	OtherAssets:        "oth_assets",                 // 其他资产
	FixedAssets:        "fix_assets",                 // 固定资产
	ConstructionMater:  "const_materials",            // 工程物资
	ConstructionProg:   "const_in_prog",              // 在建工程
	IntangibleAssets:   "intan_assets",               // 无形资产
	RightUseAssets:     "r_and_d",                    // 研发支出
	DeferredTaxAssets:  "deferred_tax_assets",        // 递延所得税资产
	CPledgeLoans:       "c_pledge_loans",             // 质押贷款
	ShortBorrows:       "st_borr",                    // 短期借款
	NotesPayable:       "notes_payable",              // 应付票据
	AccountsPayable:    "acct_payable",               // 应付账款
	AdvanceReceipts:    "adv_receipts",               // 预收款项
	ContractLiab:       "contract_liab",              // 合同负债
	EmployeeBenefits:   "empl_ben_payable",           // 应付职工薪酬
	TaxesPayable:       "taxes_payable",              // 应交税费
	TotalOtherPayable:  "tot_other_payab",            // 其他应付款合计
	EstimateLiab:       "estim_liab",                 // 预计负债
	LtBorrow:           "lt_borr",                    // 长期借款
	BondsPayable:       "bonds_payable",              // 应付债券
	LtPayable:          "lt_payable",                 // 长期应付款
	SpecialReserve:     "special_rsrv",               // 专项储备
	CapitalRece:        "cap_rese",                   // 资本公积
	SurplusReserve:     "surplus_rese",               // 盈余公积
	RetainedEarnings:   "retained_earnings",          // 未分配利润
	TreasuryShare:      "treasury_share",             // 库存股
}

// GetBalanceSheet 获取资产负债表数据
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
// - total_assets: 总资产
// - total_liab: 总负债
// - total_hldr_eqy_inc_min_int: 股东权益合计(含少数股东权益)
// - total_hldr_eqy_exc_min_int: 股东权益合计(不含少数股东权益)
// - total_cur_assets: 流动资产合计
// - total_non_cur_assets: 非流动资产合计
// - total_cur_liab: 流动负债合计
// - total_non_cur_liab: 非流动负债合计
// - 更多字段见文档
func (c *Client) GetBalanceSheet(params BalanceSheetParams, fields []string) (*types.DataFrame, error) {
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
	return c.Query("balancesheet", reqParams, fields)
}

// GetLatestBalanceSheet 获取最新的资产负债表数据（简化接口）
func (c *Client) GetLatestBalanceSheet(tsCode string) (*types.DataFrame, error) {
	return c.GetBalanceSheet(BalanceSheetParams{
		TSCode:     tsCode,
		ReportType: "1", // 默认获取合并报表
	}, nil)
}

// GetYearBalanceSheet 获取年度资产负债表数据（简化接口）
func (c *Client) GetYearBalanceSheet(tsCode string, year string) (*types.DataFrame, error) {
	return c.GetBalanceSheet(BalanceSheetParams{
		TSCode:     tsCode,
		Period:     year + "1231", // 年度报表日期
		ReportType: "1",           // 合并报表
	}, nil)
}

// GetQuarterBalanceSheet 获取季度资产负债表数据（简化接口）
func (c *Client) GetQuarterBalanceSheet(tsCode string, yearQuarter string) (*types.DataFrame, error) {
	return c.GetBalanceSheet(BalanceSheetParams{
		TSCode:     tsCode,
		Period:     yearQuarter, // 如"20211231", "20220331"
		ReportType: "1",         // 合并报表
	}, nil)
}

// CommonBalanceSheetFields 返回常用的资产负债表字段列表
func (c *Client) CommonBalanceSheetFields() []string {
	return []string{
		BalanceSheetField.TSCode,
		BalanceSheetField.EndDate,
		BalanceSheetField.AnnDate,
		BalanceSheetField.TotalAssets,
		BalanceSheetField.TotalLiab,
		BalanceSheetField.TotalSHEquity,
		BalanceSheetField.TotalCurrAssets,
		BalanceSheetField.TotalNonCurrAssets,
		BalanceSheetField.TotalCurrLiab,
		BalanceSheetField.TotalNonCurrLiab,
	}
}
