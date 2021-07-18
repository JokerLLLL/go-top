package top

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	OpenRouterHttpAPI        = "http://gw.api.taobao.com/router/rest"     //正式版API路由地址 http 版本
	OpenRouterHttpsAPI       = "https://top-proxy.internal.uco.com/router/rest" //正式版API路由地址 https 版本
	OpenSanBoxRouterHttpAPI  = "http://gw.api.tbsandbox.com/router/rest"  //沙盒版API路由地址 http 版本
	OpenSanBoxRouterHttpsAPI = "https://gw.api.tbsandbox.com/router/rest" //沙盒版API路由地址 https 版本
)

type TopClient struct {
	AppKey    string
	AppSecret string
	Session   string
	Param     map[string]interface{}
	IsHttps   bool
	IsSanBox  bool
}

func CreateTopClient(appkey, appsecret, session string) *TopClient {
	return &TopClient{
		AppKey:    appkey,
		AppSecret: appsecret,
		Session:   session,
		IsHttps:   true,
		IsSanBox:  false,
	}
}

/// TODO move to other place
var httpClient *http.Client
func HttpClient() *http.Client {
	if httpClient == nil {
		httpClient = &http.Client{}
		httpClient.Timeout = 10 * time.Second
		//连接池设置
		httpClient.Transport = &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second, // 连接超时时间
				KeepAlive: 60 * time.Second, // 保持长连接的时间
			}).DialContext, // 设置连接的参数
			MaxIdleConns:          500, // 最大空闲连接
			IdleConnTimeout:       60 * time.Second, // 空闲连接的超时时间
			ExpectContinueTimeout: 20 * time.Second, // 等待服务第一个响应的超时时间
			MaxIdleConnsPerHost:   100, // 每个host保持的空闲连接数
		}
	}
	return httpClient
}


func (t *TopClient) Run(api TopApi) ([]byte, error) {
	if msg, ok := api.CheckParam(); !ok {
		return []byte(msg), nil
	}

	hh, _ := time.ParseDuration("8h")
	loc := time.Now().UTC().Add(hh)
	timestamp := strconv.FormatInt(loc.Unix(), 10)

	param := map[string]interface{}{
		"app_key":     t.AppKey,
		"sign_method": "md5",
		"format":      "json",
		"v":           "2.0",
		"timestamp":  timestamp,
	}

	if t.Session != "" {
		param["session"] = t.Session
	}

	param["method"] = api.ApiName()
	for key, val := range api.getParam() {
		param[key] = val
	}

	param["sign"] = getSign(param, t.AppSecret)

	var routerApi string
	if t.IsSanBox {
		if t.IsHttps {
			routerApi = OpenSanBoxRouterHttpsAPI
		} else {
			routerApi = OpenSanBoxRouterHttpAPI
		}
	} else {
		if t.IsHttps {
			routerApi = OpenRouterHttpsAPI
		} else {
			routerApi = OpenRouterHttpAPI
		}
	}

	return remoteCall(routerApi, param)
}

// 获取签名
func getSign(params map[string]interface{}, appsecret string) string {
	// 获取Key
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	// 排序asc
	sort.Strings(keys)
	// 把所有参数名和参数值串在一起
	query := bytes.NewBufferString(appsecret)
	for _, k := range keys {
		query.WriteString(k)
		query.WriteString(interfaceToString(params[k]))
	}
	query.WriteString(appsecret)
	//md5
	h := md5.New()
	h.Write(query.Bytes())
	// 把二进制转化为大写的十六进制
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func remoteCall(urlStr string, param map[string]interface{}) ([]byte, error) {
	// build request
	var req *http.Request
	args := url.Values{}
	for key, val := range param {
		args.Set(key, interfaceToString(val))
	}
	paramString := args.Encode()
	//logInfo
	fmt.Println(paramString)
	fmt.Println(urlStr)

	req, err := http.NewRequest("POST", urlStr, strings.NewReader(paramString))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	// do request
	response, err := HttpClient().Do(req)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		err = fmt.Errorf("请求错误:%d", response.StatusCode)
		return nil, err
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func interfaceToString(src interface{}) string {
	if src == nil {
		return ""
	}
	switch src.(type) {
	case string:
		return src.(string)
	case int, int8, int32, int64:
	case uint8, uint16, uint32, uint64:
	case float32, float64:
		return fmt.Sprint(src)
	}
	data, err := json.Marshal(src)
	if err != nil {
		panic(err)
	}
	return string(data)
}
