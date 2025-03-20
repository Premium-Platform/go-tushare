package main

import (
	"encoding/json"
	"fmt"
	"os"

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

	// 示例1: 获取股票列表，只获取基础字段
	fmt.Println("=== 获取股票列表（基础字段）===")
	fields := []string{
		client.StockBasicField.TSCode,
		client.StockBasicField.Name,
		client.StockBasicField.Area,
		client.StockBasicField.Industry,
		client.StockBasicField.Market,
		client.StockBasicField.ListDate,
	}

	df, err := cli.GetStockBasic(client.StockBasicParams{
		ListStatus: "L", // 上市
	}, fields)

	if err != nil {
		fmt.Printf("获取数据失败: %v\n", err)
		return
	}

	// 打印前5条数据
	printDataFrameLimit(df, 5)

	// 示例2: 获取科创板股票
	fmt.Println("\n=== 获取科创板股票 ===")
	starDf, err := cli.GetSTARStocks()
	if err != nil {
		fmt.Printf("获取数据失败: %v\n", err)
		return
	}

	// 打印前5条数据
	printDataFrameLimit(starDf, 5)

	// 示例3: 获取创业板股票
	fmt.Println("\n=== 获取创业板股票 ===")
	gemDf, err := cli.GetGEMStocks()
	if err != nil {
		fmt.Printf("获取数据失败: %v\n", err)
		return
	}

	// 打印前5条数据
	printDataFrameLimit(gemDf, 5)

	// 示例4: 获取沪股通标的
	fmt.Println("\n=== 获取沪股通标的 ===")
	hkConnectDf, err := cli.GetStockBasic(client.StockBasicParams{
		ListStatus: "L",
		IsHS:       "H", // 沪股通标的
	}, cli.CommonStockFields())

	if err != nil {
		fmt.Printf("获取数据失败: %v\n", err)
		return
	}

	// 打印前5条数据
	printDataFrameLimit(hkConnectDf, 5)
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
