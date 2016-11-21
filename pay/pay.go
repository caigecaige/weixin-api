// pay.go 微信支付
package pay

import (
	"crypto/md5"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	PREPAY_API = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	QUERY_API  = "https://api.mch.weixin.qq.com/pay/orderquery"
)
const (
	WX_ORDERS_TRADE_TYPE  = "NAVITE"
	WX_ORDERS_DEVICE_INFO = "bcjishu"
)

const (
	TRADE_TYPE_JSAPI  = "JSAPI"
	TRADE_TYPE_NATIVE = "NATIVE"
	TRADE_TYPE_APP    = "APP"
)

const SUCCESS = "SUCCESS"
const WX_API_NOTIFY_URL = "/callback"

//商户帐号
type WxPayAccount struct {
	AppId          string `json:"appid"`
	MerchantId     string `json:"merchantId"`
	MerchantApiKey string `json:"merchantApiKey"`
	AccName        string `json:"accName"`
	AccKey         string `json:"accKey"`
}

//微信提交请求
type WxPostRequest struct {
	XMLName        xml.Name `xml:"xml"`
	AppId          string   `xml:"appid"`
	MchId          string   `xml:"mch_id"`
	DeviceInfo     string   `xml:"device_info"`
	NonceStr       string   `xml:"nonce_str"`
	Body           string   `xml:"body"`
	Detail         string   `xml:"detail"`
	Attach         string   `xml:"attach"`
	OutTradeNo     string   `xml:"out_trade_no"`
	FeeType        string   `xml:"fee_type"`
	TotalFee       int      `xml:"total_fee"` //整型,单位为分
	SpbillCreateIp string   `xml:"spbill_create_ip"`
	TimeStart      string   `xml:"time_start"`
	TimeExpire     string   `xml:"time_expire"`
	GoodsTag       string   `xml:"goods_tag"`
	NotifyUrl      string   `xml:"notify_url"`
	TradeType      string   `xml:"trade_type"`
	ProductId      string   `xml:"product_id"`
	OpenId         string   `xml:"openid"`
	Sign           string   `xml:"sign"`
}

//微信返回
type WxPostResponse struct {
	XMLName    xml.Name `xml:"xml"`
	ReturnCode string   `xml:"return_code"`
	ReturnMsg  string   `xml:"return_msg"`
	AppId      string   `xml:"appid"`
	MchId      string   `xml:"mch_id"`
	DeviceInfo string   `xml:"device_info"`
	NonceStr   string   `xml:"nonce_str"`
	Sign       string   `xml:"sign"`
	ResultCode string   `xml:"result_code"`
	ErrCode    string   `xml:"err_code"`
	ErrCodeDes string   `xml:"err_code_des"`
	TradeType  string   `xml:"trade_type"`
	PrepayId   string   `xml:"prepay_id"`
	CodeUrl    string   `xml:"code_url"`
}

//微信查询请求
type WxQueryRequest struct {
	XMLName       xml.Name `xml:"xml"`
	AppId         string   `xml:"appid"`
	MchId         string   `xml:"mch_id"`
	TransactionId string   `xml:"transaction_id"`
	OutTradeNo    string   `xml:"out_trade_no"`
	NonceStr      string   `xml:"nonce_str"`
	Sign          string   `xml:"sign"`
}

//查询返回
type WxQueryResponse struct {
	XMLName        xml.Name `xml:"xml"`
	ReturnCode     string   `xml:"return_code"`
	ReturnMsg      string   `xml:"return_msg"`
	ResultCode     string   `xml:"result_code"`
	ErrCode        string   `xml:"err_code"`
	ErrCode_des    string   `xml:"err_code_des"`
	DeviceInfo     string   `xml:"device_info"`
	Openid         string   `xml:"openid"`
	IsSubscribe    string   `xml:"is_subscribe"`
	TradeType      string   `xml:"trade_type"`
	TradeState     string   `xml:"trade_state"`
	BankType       string   `xml:"bank_type"`
	TotalFee       string   `xml:"total_fee"`
	FeeType        string   `xml:"fee_type"`
	CashFee        string   `xml:"cash_fee"`
	CashFeeType    string   `xml:"cash_fee_type"`
	TransactionId  string   `xml:"transaction_id"`
	OutTradeNo     string   `xml:"out_trade_no"`
	Attach         string   `xml:"attach"`
	TimeEnd        string   `xml:"time_end"`
	PayTime        int64
	TradeStateDesc string `xml:"trade_state_desc"`
}

//微信支付
type WxPay struct {
	url           string
	Account       WxPayAccount
	PostRequest   WxPostRequest
	PostResponse  WxPostResponse
	QueryRequest  WxQueryRequest
	QueryResponse WxQueryResponse
}

//初始化
//生成签名
func (wp *WxPay) initPost() {
	wp.url = PREPAY_API
	wp.PostRequest.AppId = wp.Account.AppId
	wp.PostRequest.MchId = wp.Account.MerchantId
	var params []map[string]string
	params = append(params, map[string]string{"key": "appid", "val": wp.PostRequest.AppId})
	params = append(params, map[string]string{"key": "attach", "val": wp.PostRequest.Attach})
	params = append(params, map[string]string{"key": "body", "val": wp.PostRequest.Body})
	params = append(params, map[string]string{"key": "detail", "val": wp.PostRequest.Detail})
	params = append(params, map[string]string{"key": "device_info", "val": wp.PostRequest.DeviceInfo})
	params = append(params, map[string]string{"key": "fee_type", "val": wp.PostRequest.FeeType})
	params = append(params, map[string]string{"key": "goods_tag", "val": wp.PostRequest.GoodsTag})
	params = append(params, map[string]string{"key": "mch_id", "val": wp.PostRequest.MchId})
	params = append(params, map[string]string{"key": "nonce_str", "val": wp.PostRequest.NonceStr})
	params = append(params, map[string]string{"key": "notify_url", "val": wp.PostRequest.NotifyUrl})
	params = append(params, map[string]string{"key": "openid", "val": wp.PostRequest.OpenId})
	params = append(params, map[string]string{"key": "out_trade_no", "val": wp.PostRequest.OutTradeNo})
	params = append(params, map[string]string{"key": "product_id", "val": wp.PostRequest.ProductId})
	params = append(params, map[string]string{"key": "spbill_create_ip", "val": wp.PostRequest.SpbillCreateIp})
	params = append(params, map[string]string{"key": "time_start", "val": wp.PostRequest.TimeStart})
	params = append(params, map[string]string{"key": "time_expire", "val": wp.PostRequest.TimeExpire})
	params = append(params, map[string]string{"key": "total_fee", "val": strconv.Itoa(wp.PostRequest.TotalFee)})
	params = append(params, map[string]string{"key": "trade_type", "val": wp.PostRequest.TradeType})
	params = append(params, map[string]string{"key": "key", "val": wp.Account.MerchantApiKey})
	sign := GenerateSign(params)
	wp.PostRequest.Sign = sign
}

//提交
func (wp *WxPay) Post() bool {
	wp.initPost()
	data, _ := xml.Marshal(wp.PostRequest)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //不检验证书
	}
	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest("POST", wp.url, strings.NewReader(string(data)))
	req.Header.Set("Accept", "text/xml,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Charset", "UTF-8")
	res, err := client.Do(req)
	if err != nil {
		fmt.Sprintln(err)
		panic(err)
	}
	defer res.Body.Close()
	content, rerr := ioutil.ReadAll(res.Body)
	if rerr != nil {
		return false
	}
	xml.Unmarshal([]byte(content), &wp.PostResponse)
	return true

}

//初始化查询
func (wp *WxPay) initQuery() {
	wp.url = QUERY_API
	//生成查询的签名
	var params []map[string]string
	params = append(params, map[string]string{"key": "appid", "val": wp.Account.AppId})
	params = append(params, map[string]string{"key": "mch_id", "val": wp.Account.MerchantId})
	params = append(params, map[string]string{"key": "nonce_str", "val": wp.QueryRequest.NonceStr})
	params = append(params, map[string]string{"key": "out_trade_no", "val": wp.QueryRequest.OutTradeNo})
	params = append(params, map[string]string{"key": "key", "val": wp.Account.MerchantApiKey})
	wp.QueryRequest.AppId = wp.Account.AppId
	wp.QueryRequest.MchId = wp.Account.MerchantId
	wp.QueryRequest.Sign = GenerateSign(params)
}

//查询
func (wp *WxPay) Query() {
	wp.initQuery()
	data, _ := xml.Marshal(wp.QueryRequest)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //不检验证书
	}
	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest("POST", wp.url, strings.NewReader(string(data)))
	req.Header.Set("Accept", "text/xml,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Charset", "UTF-8")
	res, err := client.Do(req)
	if err != nil {
		fmt.Sprintln(err)
		panic(err)
	}
	defer res.Body.Close()
	content, _ := ioutil.ReadAll(res.Body)
	xml.Unmarshal([]byte(content), &wp.QueryResponse)
	payTime, timeErr := time.ParseInLocation("20060102150405", wp.QueryResponse.TimeEnd, time.Local)
	if timeErr == nil {
		wp.QueryResponse.PayTime = payTime.Unix()
	}
	return
}

//生成签名
func GenerateSign(params []map[string]string) string {
	var paramStr string
	l := len(params)
	for i := 0; i < l; i++ {
		if params[i]["val"] == "" {
			continue
		}
		if i == (l - 1) {
			paramStr = paramStr + params[i]["key"] + "=" + params[i]["val"]
		} else {
			paramStr = paramStr + params[i]["key"] + "=" + params[i]["val"] + "&"
		}
	}
	signChar := md5.Sum([]byte(paramStr))
	sign := fmt.Sprintf("%x", signChar)
	return sign
}
