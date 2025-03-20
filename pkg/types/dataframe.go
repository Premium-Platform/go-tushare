package types

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
)

// DataFrame 是数据结果的通用结构
type DataFrame struct {
	Columns []string                 `json:"columns"`
	Rows    []map[string]interface{} `json:"rows"`
}

// NewDataFrame 创建一个新的DataFrame
func NewDataFrame(columns []string, rows []map[string]interface{}) *DataFrame {
	return &DataFrame{
		Columns: columns,
		Rows:    rows,
	}
}

// GetColumn 获取指定列的数据
func (df *DataFrame) GetColumn(name string) []interface{} {
	result := make([]interface{}, 0, len(df.Rows))
	for _, row := range df.Rows {
		if val, ok := row[name]; ok {
			result = append(result, val)
		} else {
			result = append(result, nil)
		}
	}
	return result
}

// ToJSON 将DataFrame转换为JSON
func (df *DataFrame) ToJSON() ([]byte, error) {
	return json.Marshal(df)
}

// ToCSV 将DataFrame转换为CSV
func (df *DataFrame) ToCSV() ([]byte, error) {
	buffer := &bytes.Buffer{}
	writer := csv.NewWriter(buffer)

	// 写入表头
	if err := writer.Write(df.Columns); err != nil {
		return nil, err
	}

	// 写入数据行
	for _, row := range df.Rows {
		record := make([]string, len(df.Columns))
		for i, col := range df.Columns {
			if val, ok := row[col]; ok {
				record[i] = toString(val)
			}
		}
		if err := writer.Write(record); err != nil {
			return nil, err
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

// toString 将任意类型转换为字符串
func toString(v interface{}) string {
	if v == nil {
		return ""
	}

	switch val := v.(type) {
	case string:
		return val
	case []byte:
		return string(val)
	default:
		b, _ := json.Marshal(val)
		return string(b)
	}
}
