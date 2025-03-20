package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	tsError "github.com/Premium-Platform/go-tushare/pkg/errors"
	"github.com/Premium-Platform/go-tushare/pkg/logger"
	"github.com/Premium-Platform/go-tushare/pkg/types"
)

const (
	// DefaultTimeout 默认超时时间
	DefaultTimeout = 30 * time.Second

	// DefaultAPIURL 默认API地址
	DefaultAPIURL = "http://api.tushare.pro"
)

// Client TuShare API客户端
type Client struct {
	token   string
	apiURL  string
	timeout time.Duration
	client  *http.Client
	logger  *logger.Logger
}

// RequestParams 请求参数
type RequestParams struct {
	APIName string                 `json:"api_name"`
	Token   string                 `json:"token"`
	Params  map[string]interface{} `json:"params"`
	Fields  []string               `json:"fields"`
}

// ResponseData API响应数据
type ResponseData struct {
	Code int `json:"code"`
	Data struct {
		Fields []string        `json:"fields"`
		Items  [][]interface{} `json:"items"`
	} `json:"data"`
	Message string `json:"msg"`
}

// New 创建一个新的客户端
func New(token string) *Client {
	client := &Client{
		token:   token,
		apiURL:  DefaultAPIURL,
		timeout: DefaultTimeout,
		client:  &http.Client{Timeout: DefaultTimeout},
		logger:  logger.NewLogger(nil, logger.INFO),
	}
	return client
}

// SetToken 设置令牌
func (c *Client) SetToken(token string) {
	c.token = token
	c.logger.Info("Token已更新")
}

// GetToken 获取令牌
func (c *Client) GetToken() string {
	return c.token
}

// SetTimeout 设置超时时间
func (c *Client) SetTimeout(timeout time.Duration) {
	c.timeout = timeout
	c.client.Timeout = timeout
	c.logger.Info("超时时间已设置为 %v", timeout)
}

// SetAPIURL 设置API地址
func (c *Client) SetAPIURL(url string) {
	c.apiURL = url
	c.logger.Info("API地址已设置为 %s", url)
}

// SetLogger 设置日志记录器
func (c *Client) SetLogger(l *logger.Logger) {
	c.logger = l
}

// Query 通用API查询
func (c *Client) Query(apiName string, params map[string]interface{}, fields []string) (*types.DataFrame, error) {
	// 检查token
	if c.token == "" {
		c.logger.Error("无效的Token")
		return nil, tsError.ErrInvalidToken
	}

	c.logger.Debug("开始查询API: %s, 参数: %v", apiName, params)

	// 构建请求参数
	reqParams := RequestParams{
		APIName: apiName,
		Token:   c.token,
		Params:  params,
		Fields:  fields,
	}

	// 转换为JSON
	reqData, err := json.Marshal(reqParams)
	if err != nil {
		c.logger.Error("请求参数序列化失败: %v", err)
		return nil, tsError.Wrap(err, "failed to marshal request params")
	}

	// 创建请求
	req, err := http.NewRequest(http.MethodPost, c.apiURL, bytes.NewBuffer(reqData))
	if err != nil {
		c.logger.Error("创建HTTP请求失败: %v", err)
		return nil, tsError.Wrap(err, "failed to create request")
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	c.logger.Debug("发送请求到 %s", c.apiURL)
	resp, err := c.client.Do(req)
	if err != nil {
		c.logger.Error("发送请求失败: %v", err)
		return nil, tsError.Wrap(err, "failed to send request")
	}
	defer resp.Body.Close()

	// 读取响应数据
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error("读取响应数据失败: %v", err)
		return nil, tsError.Wrap(err, "failed to read response body")
	}

	// 解析响应数据
	var respObj ResponseData
	if err := json.Unmarshal(respData, &respObj); err != nil {
		c.logger.Error("解析响应数据失败: %v", err)
		return nil, tsError.Wrap(err, "failed to unmarshal response data")
	}

	// 检查响应状态
	if respObj.Code != 0 {
		c.logger.Error("API返回错误: 代码=%d, 消息=%s", respObj.Code, respObj.Message)
		return nil, tsError.NewAPIError(respObj.Code, respObj.Message)
	}

	// 将响应数据转换为DataFrame
	columns := respObj.Data.Fields
	rows := make([]map[string]interface{}, len(respObj.Data.Items))

	for i, item := range respObj.Data.Items {
		row := make(map[string]interface{})
		for j, field := range columns {
			if j < len(item) {
				row[field] = item[j]
			}
		}
		rows[i] = row
	}

	c.logger.Debug("查询成功, 返回 %d 行数据", len(rows))
	return types.NewDataFrame(columns, rows), nil
}
