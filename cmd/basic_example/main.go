package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/Premium-Platform/go-tushare/client"
	"github.com/Premium-Platform/go-tushare/pkg/types"
)

func main() {
	// 从环境变量获取token
	token := os.Getenv("TUSHARE_TOKEN")
	if token == "" {
		fmt.Println("请设置环境变量TUSHARE_TOKEN")
		return
	}

	// 创建客户端
	cli := client.New(token)

	// 示例1: 获取交易日历
	fmt.Println("=== 获取交易日历 ===")
	now := time.Now()
	startDate := now.AddDate(0, -1, 0).Format("20060102") // 一个月前
	endDate := now.Format("20060102")                     // 今天

	calDf, err := cli.GetTradeDays(startDate, endDate)
	if err != nil {
		fmt.Printf("获取交易日历数据失败: %v\n", err)
	} else {
		printDataFrameLimit(calDf, 5)
	}

	// 示例2: 获取股票曾用名
	fmt.Println("\n=== 获取股票曾用名（以平安银行为例） ===")
	nameChangeDf, err := cli.GetStockNameHistory("000001.SZ")
	if err != nil {
		fmt.Printf("获取股票曾用名数据失败: %v\n", err)
	} else {
		printDataFrameLimit(nameChangeDf, 5)
	}

	// 示例3: 获取沪深股通成份股
	fmt.Println("\n=== 获取沪股通成份股 ===")
	hsConstDf, err := cli.GetSHConst()
	if err != nil {
		fmt.Printf("获取沪股通成份股数据失败: %v\n", err)
	} else {
		printDataFrameLimit(hsConstDf, 5)
	}

	// 示例4: 获取上市公司基本信息
	fmt.Println("\n=== 获取上市公司基本信息（以平安银行为例） ===")
	companyDf, err := cli.GetCompanyInfo("000001.SZ")
	if err != nil {
		fmt.Printf("获取上市公司基本信息失败: %v\n", err)
	} else {
		printDataFrameLimit(companyDf, 1)
	}

	// 示例5: 获取新股上市信息
	fmt.Println("\n=== 获取近期新股上市信息 ===")
	newShareDf, err := cli.GetRecentNewShares()
	if err != nil {
		fmt.Printf("获取新股上市信息失败: %v\n", err)
	} else {
		printDataFrameLimit(newShareDf, 5)
	}
}

// 打印DataFrame数据（有限条数）
func printDataFrameLimit(df *types.DataFrame, limit int) {
	if df == nil || len(df.Rows) == 0 {
		fmt.Println("无数据")
		return
	}

	if limit > len(df.Rows) {
		limit = len(df.Rows)
	}

	limitedDf := &types.DataFrame{
		Columns: df.Columns,
		Rows:    df.Rows[:limit],
	}

	data, err := json.MarshalIndent(limitedDf, "", "  ")
	if err != nil {
		fmt.Printf("序列化数据失败: %v\n", err)
		return
	}

	fmt.Println(string(data))
	fmt.Printf("显示了 %d/%d 条记录\n", limit, len(df.Rows))
}
