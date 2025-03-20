/*
 * @Author: SpenserCai
 * @Date: 2025-03-20 21:31:51
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2025-03-21 00:04:03
 * @Description: file content
 */
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Premium-Platform/go-tushare/client"
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

	// 示例1: 获取股票日线行情数据
	fmt.Println("=== 获取股票日线行情数据 ===")
	df, err := cli.Query("daily", map[string]interface{}{
		"ts_code":    "000001.SZ",
		"start_date": "20220101",
		"end_date":   "20220110",
	}, []string{})

	if err != nil {
		fmt.Printf("获取数据失败: %v\n", err)
		return
	}

	// 打印结果
	printDataFrame(df)

	// 示例2: 使用通用行情接口获取数据
	fmt.Println("\n=== 使用通用行情接口获取数据 ===")
	barDf, err := cli.Bar(client.BarParams{
		TsCode:     "000001.SZ",
		StartDate:  "20220101",
		EndDate:    "20220110",
		Freq:       "D",
		AssetType:  "E",
		AdjustType: "qfq",
	})

	if err != nil {
		fmt.Printf("获取数据失败: %v\n", err)
		return
	}

	// 打印结果
	printDataFrame(barDf)
}

// 打印DataFrame数据
func printDataFrame(df interface{}) {
	data, err := json.MarshalIndent(df, "", "  ")
	if err != nil {
		fmt.Printf("序列化数据失败: %v\n", err)
		return
	}

	fmt.Println(string(data))
}
