package client

import (
	"github.com/Premium-Platform/go-tushare/pkg/types"
)

// StockCompanyParams 上市公司基本信息查询参数
type StockCompanyParams struct {
	TsCode   string `json:"ts_code"`  // 股票代码
	Exchange string `json:"exchange"` // 交易所代码，SSE上交所 SZSE深交所
}

// StockCompanyField 上市公司基本信息字段常量
var StockCompanyField = struct {
	TsCode        string
	Exchange      string
	Chairman      string
	Manager       string
	Secretary     string
	RegCapital    string
	SetupDate     string
	Province      string
	City          string
	Introduction  string
	Website       string
	Email         string
	Office        string
	Employees     string
	MainBusiness  string
	BusinessScope string
}{
	TsCode:        "ts_code",        // 股票代码
	Exchange:      "exchange",       // 交易所代码
	Chairman:      "chairman",       // 法人代表
	Manager:       "manager",        // 总经理
	Secretary:     "secretary",      // 董秘
	RegCapital:    "reg_capital",    // 注册资本
	SetupDate:     "setup_date",     // 注册日期
	Province:      "province",       // 所在省份
	City:          "city",           // 所在城市
	Introduction:  "introduction",   // 公司介绍
	Website:       "website",        // 公司主页
	Email:         "email",          // 电子邮件
	Office:        "office",         // 办公室
	Employees:     "employees",      // 员工人数
	MainBusiness:  "main_business",  // 主要业务及产品
	BusinessScope: "business_scope", // 经营范围
}

// GetStockCompany 获取上市公司基本信息
//
// 接口参数：
// - ts_code: 股票代码
// - exchange: 交易所代码，SSE上交所 SZSE深交所
//
// 返回字段：
// - ts_code: 股票代码
// - exchange: 交易所代码
// - chairman: 法人代表
// - manager: 总经理
// - secretary: 董秘
// - reg_capital: 注册资本
// - setup_date: 注册日期
// - province: 所在省份
// - city: 所在城市
// - introduction: 公司介绍
// - website: 公司主页
// - email: 电子邮件
// - office: 办公室
// - employees: 员工人数
// - main_business: 主要业务及产品
// - business_scope: 经营范围
func (c *Client) GetStockCompany(params StockCompanyParams, fields []string) (*types.DataFrame, error) {
	// 构建请求参数
	reqParams := map[string]interface{}{}

	if params.TsCode != "" {
		reqParams["ts_code"] = params.TsCode
	}

	if params.Exchange != "" {
		reqParams["exchange"] = params.Exchange
	}

	// 调用通用查询接口
	return c.Query("stock_company", reqParams, fields)
}

// GetCompanyInfo 获取单个公司基本信息（简化接口）
func (c *Client) GetCompanyInfo(tsCode string) (*types.DataFrame, error) {
	return c.GetStockCompany(StockCompanyParams{
		TsCode: tsCode,
	}, []string{})
}

// GetSSECompanies 获取上交所上市公司
func (c *Client) GetSSECompanies() (*types.DataFrame, error) {
	return c.GetStockCompany(StockCompanyParams{
		Exchange: "SSE",
	}, []string{})
}

// GetSZSECompanies 获取深交所上市公司
func (c *Client) GetSZSECompanies() (*types.DataFrame, error) {
	return c.GetStockCompany(StockCompanyParams{
		Exchange: "SZSE",
	}, []string{})
}

// CommonStockCompanyFields 返回上市公司基本信息常用字段
func (c *Client) CommonStockCompanyFields() []string {
	return []string{
		StockCompanyField.TsCode,
		StockCompanyField.Exchange,
		StockCompanyField.Chairman,
		StockCompanyField.Manager,
		StockCompanyField.Secretary,
		StockCompanyField.Province,
		StockCompanyField.City,
		StockCompanyField.Website,
		StockCompanyField.Email,
		StockCompanyField.MainBusiness,
	}
}
