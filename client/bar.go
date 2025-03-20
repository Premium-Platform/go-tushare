package client

import (
	"fmt"

	"github.com/Premium-Platform/go-tushare/pkg/types"
)

// BarParams Bar接口参数
type BarParams struct {
	TsCode       string   // 证券代码
	StartDate    string   // 开始日期 YYYYMMDD
	EndDate      string   // 结束日期 YYYYMMDD
	Freq         string   // 周期：D=日线 W=周线 M=月线
	AssetType    string   // 资产类别：E=股票 I=指数 C=数字货币 FT=期货
	Exchange     string   // 交易所
	AdjustType   string   // 复权类型：None=不复权 qfq=前复权 hfq=后复权
	MA           []int    // 均线
	Factors      []string // 因子数据
	ContractType string   // 合约类型
}

// FreqMap 频率映射
var FreqMap = map[string]string{
	"D": "1DAY",
	"W": "1WEEK",
	"M": "1MONTH",
	"Y": "1YEAR",
}

// Bar 行情数据通用接口
func (c *Client) Bar(params BarParams) (*types.DataFrame, error) {
	// 股票
	if params.AssetType == "E" || params.AssetType == "" {
		return c.stockBar(params)
	}

	// 指数
	if params.AssetType == "I" {
		return c.indexBar(params)
	}

	// 期货
	if params.AssetType == "FT" {
		return c.futureBar(params)
	}

	// 数字货币
	if params.AssetType == "C" {
		return c.coinBar(params)
	}

	// 默认使用股票接口
	return c.stockBar(params)
}

// stockBar 股票行情数据
func (c *Client) stockBar(params BarParams) (*types.DataFrame, error) {
	var apiName string

	// 根据周期选择API
	switch params.Freq {
	case "D":
		apiName = "daily"
	case "W":
		apiName = "weekly"
	case "M":
		apiName = "monthly"
	default:
		apiName = "daily"
	}

	// 构建请求参数
	queryParams := map[string]interface{}{
		"ts_code":    params.TsCode,
		"start_date": params.StartDate,
		"end_date":   params.EndDate,
	}

	// 获取行情数据
	df, err := c.Query(apiName, queryParams, []string{})
	if err != nil {
		return nil, err
	}

	// 如果需要复权处理
	if params.AdjustType != "" && params.AdjustType != "None" {
		df, err = c.adjustBar(df, params)
		if err != nil {
			return nil, err
		}
	}

	// 如果需要计算因子
	if len(params.Factors) > 0 {
		df, err = c.addFactors(df, params)
		if err != nil {
			return nil, err
		}
	}

	// 如果需要计算均线
	if len(params.MA) > 0 {
		df, err = c.calculateMA(df, params.MA)
		if err != nil {
			return nil, err
		}
	}

	return df, nil
}

// indexBar 指数行情数据
func (c *Client) indexBar(params BarParams) (*types.DataFrame, error) {
	var apiName string

	// 根据周期选择API
	switch params.Freq {
	case "D":
		apiName = "index_daily"
	case "W":
		apiName = "index_weekly"
	case "M":
		apiName = "index_monthly"
	default:
		apiName = "index_daily"
	}

	// 构建请求参数
	queryParams := map[string]interface{}{
		"ts_code":    params.TsCode,
		"start_date": params.StartDate,
		"end_date":   params.EndDate,
	}

	// 获取行情数据
	df, err := c.Query(apiName, queryParams, []string{})
	if err != nil {
		return nil, err
	}

	// 如果需要计算均线
	if len(params.MA) > 0 {
		df, err = c.calculateMA(df, params.MA)
		if err != nil {
			return nil, err
		}
	}

	return df, nil
}

// futureBar 期货行情数据
func (c *Client) futureBar(params BarParams) (*types.DataFrame, error) {
	// 构建请求参数
	queryParams := map[string]interface{}{
		"ts_code":    params.TsCode,
		"start_date": params.StartDate,
		"end_date":   params.EndDate,
		"exchange":   params.Exchange,
	}

	// 获取行情数据
	df, err := c.Query("fut_daily", queryParams, []string{})
	if err != nil {
		return nil, err
	}

	// 如果需要计算均线
	if len(params.MA) > 0 {
		df, err = c.calculateMA(df, params.MA)
		if err != nil {
			return nil, err
		}
	}

	return df, nil
}

// coinBar 数字货币行情数据
func (c *Client) coinBar(params BarParams) (*types.DataFrame, error) {
	// 转换频率格式
	freq := "daily"
	if params.Freq == "W" {
		freq = "week"
	}

	// 构建请求参数
	queryParams := map[string]interface{}{
		"symbol":        params.TsCode,
		"start_date":    params.StartDate,
		"end_date":      params.EndDate,
		"exchange":      params.Exchange,
		"freq":          freq,
		"contract_type": params.ContractType,
	}

	// 获取行情数据
	df, err := c.Query("coinbar", queryParams, []string{})
	if err != nil {
		return nil, err
	}

	// 如果需要计算均线
	if len(params.MA) > 0 {
		df, err = c.calculateMA(df, params.MA)
		if err != nil {
			return nil, err
		}
	}

	return df, nil
}

// adjustBar 复权处理
func (c *Client) adjustBar(df *types.DataFrame, params BarParams) (*types.DataFrame, error) {
	if df == nil || len(df.Rows) == 0 {
		return df, nil
	}

	// 获取复权因子
	c.logger.Debug("正在获取复权因子, ts_code=%s, start_date=%s, end_date=%s",
		params.TsCode, params.StartDate, params.EndDate)

	fcts, err := c.GetAdjFactor(AdjFactorParams{
		TSCode:    params.TsCode,
		StartDate: params.StartDate,
		EndDate:   params.EndDate,
	}, nil)

	if err != nil {
		c.logger.Error("获取复权因子失败: %v", err)
		return nil, err
	}

	if fcts == nil || len(fcts.Rows) == 0 {
		c.logger.Warn("未找到复权因子数据, 将使用未复权数据")
		return df, nil
	}

	// 创建日期到复权因子的映射
	adjFactorMap := make(map[string]float64)
	var lastFactor float64 = 1.0

	for _, row := range fcts.Rows {
		date, ok := row["trade_date"].(string)
		if !ok {
			continue
		}

		factor, ok := row["adj_factor"].(float64)
		if !ok {
			// 尝试从字符串转换
			if factorStr, ok := row["adj_factor"].(string); ok {
				if f, err := stringToFloat(factorStr); err == nil {
					factor = f
				} else {
					continue
				}
			} else {
				continue
			}
		}

		adjFactorMap[date] = factor
		lastFactor = factor
	}

	// 如果没有复权因子数据，则返回原始数据
	if len(adjFactorMap) == 0 {
		c.logger.Warn("复权因子数据为空, 将使用未复权数据")
		return df, nil
	}

	// 获取第一个交易日的复权因子（用于前复权）
	var firstFactor float64 = 0
	for _, row := range fcts.Rows {
		if factor, ok := row["adj_factor"].(float64); ok {
			firstFactor = factor
			break
		} else if factorStr, ok := row["adj_factor"].(string); ok {
			if f, err := stringToFloat(factorStr); err == nil {
				firstFactor = f
				break
			}
		}
	}

	if firstFactor == 0 {
		firstFactor = 1.0
	}

	// 对每一行数据进行复权处理
	for i, row := range df.Rows {
		date, ok := row["trade_date"].(string)
		if !ok {
			continue
		}

		factor, ok := adjFactorMap[date]
		if !ok {
			// 如果当前日期没有复权因子，使用最近的复权因子
			factor = lastFactor
		}

		// 复权处理
		if params.AdjustType == "qfq" { // 前复权
			adjFactor := factor / firstFactor
			df.Rows[i] = applyAdjustFactor(row, adjFactor)
		} else if params.AdjustType == "hfq" { // 后复权
			df.Rows[i] = applyAdjustFactor(row, factor)
		}
	}

	c.logger.Debug("复权处理完成, 处理类型: %s", params.AdjustType)
	return df, nil
}

// applyAdjustFactor 应用复权因子到行情数据
func applyAdjustFactor(row map[string]interface{}, factor float64) map[string]interface{} {
	// 复制原始数据
	result := make(map[string]interface{})
	for k, v := range row {
		result[k] = v
	}

	// 价格字段列表
	priceFields := []string{"open", "high", "low", "close", "pre_close"}

	// 对价格字段应用复权因子
	for _, field := range priceFields {
		if value, ok := row[field]; ok {
			switch v := value.(type) {
			case float64:
				result[field] = v * factor
			case string:
				if price, err := stringToFloat(v); err == nil {
					result[field] = price * factor
				}
			}
		}
	}

	return result
}

// stringToFloat 将字符串转换为浮点数
func stringToFloat(s string) (float64, error) {
	var f float64
	_, err := fmt.Sscanf(s, "%f", &f)
	return f, err
}

// calculateMA 计算均线
func (c *Client) calculateMA(df *types.DataFrame, ma []int) (*types.DataFrame, error) {
	if df == nil || len(df.Rows) == 0 || len(ma) == 0 {
		return df, nil
	}

	c.logger.Debug("开始计算均线, 周期: %v", ma)

	// 提取收盘价数据
	closes := make([]float64, len(df.Rows))
	volumes := make([]float64, len(df.Rows))

	for i, row := range df.Rows {
		// 获取收盘价
		if close, ok := row["close"].(float64); ok {
			closes[i] = close
		} else if closeStr, ok := row["close"].(string); ok {
			if c, err := stringToFloat(closeStr); err == nil {
				closes[i] = c
			}
		}

		// 获取成交量
		if vol, ok := row["vol"].(float64); ok {
			volumes[i] = vol
		} else if volStr, ok := row["vol"].(string); ok {
			if v, err := stringToFloat(volStr); err == nil {
				volumes[i] = v
			}
		}
	}

	// 计算各周期均线
	for _, period := range ma {
		if period <= 0 {
			continue
		}

		// 计算价格均线
		maField := fmt.Sprintf("ma%d", period)
		for i := range df.Rows {
			if i >= period-1 {
				sum := 0.0
				for j := 0; j < period; j++ {
					sum += closes[i-j]
				}
				df.Rows[i][maField] = sum / float64(period)
			} else {
				// 数据不足一个周期时，设为当前收盘价
				df.Rows[i][maField] = closes[i]
			}
		}

		// 计算成交量均线
		volMaField := fmt.Sprintf("vol_ma%d", period)
		for i := range df.Rows {
			if i >= period-1 {
				sum := 0.0
				for j := 0; j < period; j++ {
					sum += volumes[i-j]
				}
				df.Rows[i][volMaField] = sum / float64(period)
			} else {
				// 数据不足一个周期时，设为当前成交量
				df.Rows[i][volMaField] = volumes[i]
			}
		}
	}

	c.logger.Debug("均线计算完成")
	return df, nil
}

// addFactors 添加因子数据
func (c *Client) addFactors(df *types.DataFrame, params BarParams) (*types.DataFrame, error) {
	if df == nil || len(df.Rows) == 0 || len(params.Factors) == 0 {
		return df, nil
	}

	c.logger.Debug("开始计算因子数据, 因子: %v", params.Factors)

	// 处理不同的因子
	for _, factor := range params.Factors {
		switch factor {
		case "vr", "volume_ratio":
			df = c.calculateVolumeRatio(df)
		case "tor", "turnover_rate":
			df = c.calculateTurnoverRate(df, params.TsCode)
		}
	}

	c.logger.Debug("因子数据计算完成")
	return df, nil
}

// calculateVolumeRatio 计算量比
func (c *Client) calculateVolumeRatio(df *types.DataFrame) *types.DataFrame {
	if len(df.Rows) < 6 {
		// 数据不足，无法计算量比
		return df
	}

	// 计算前5日平均成交量
	for i := 5; i < len(df.Rows); i++ {
		var sum float64
		for j := 1; j <= 5; j++ {
			vol := 0.0
			if v, ok := df.Rows[i-j]["vol"].(float64); ok {
				vol = v
			} else if vStr, ok := df.Rows[i-j]["vol"].(string); ok {
				if v, err := stringToFloat(vStr); err == nil {
					vol = v
				}
			}
			sum += vol
		}

		avgVol := sum / 5.0

		// 当日成交量
		curVol := 0.0
		if v, ok := df.Rows[i]["vol"].(float64); ok {
			curVol = v
		} else if vStr, ok := df.Rows[i]["vol"].(string); ok {
			if v, err := stringToFloat(vStr); err == nil {
				curVol = v
			}
		}

		// 计算量比
		if avgVol > 0 {
			df.Rows[i]["volume_ratio"] = curVol / avgVol
		} else {
			df.Rows[i]["volume_ratio"] = 0
		}
	}

	return df
}

// calculateTurnoverRate 计算换手率
func (c *Client) calculateTurnoverRate(df *types.DataFrame, tsCode string) *types.DataFrame {
	// 此处需要获取股票的总股本数据来计算换手率
	// 由于这需要额外的API调用，暂时不实现
	// TODO: 实现换手率计算

	return df
}
