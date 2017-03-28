// main.go,商户转账给个人
package transfer

import (
	"crypto/md5"
	"crypto/tls"
	"crypto/x509"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	API   = "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers"
	Q_API = "https://api.mch.weixin.qq.com/mmpaymkttransfers/gettransferinfo"
)

const (
	CHECK_NAME_OPTION_NO     = "NO_CHECK"
	CHECK_NAME_OPTION_FORCE  = "FORCE_CHECK"
	CHECK_NAME_OPTION_OPTION = "OPTION_CHECK"
)

const (
	RETURN_CODE_SUCC = "SUCCESS"
	RETURN_CODE_FAIL = "FAIL"
	RESULT_CODE_SUCC = "SUCCESS"
	RESULT_CODE_FAIL = "FAIL"
)

const (
	TRANSFER_STATUS_SUCC    = "SUCCESS"
	TRANSFER_STATUS_FAIL    = "FAILED"
	TRANSFER_STATUS_PROCESS = "PROCESSING"
)

//判断接口是否可用
func init() {

}

//证书
type Pem struct {
	Cert string
	Key  string
	Ca   string
}

//转账请求
type PostRequest struct {
	XMLName        xml.Name `xml:"xml"`
	Amount         int      `xml:"amount"`
	CheckName      string   `xml:"check_name"`
	Desc           string   `xml:"desc"`
	DeviceInfo     string   `xml:"device_info"`
	Key            string   `xml:"-"`
	MchAppId       string   `xml:"mch_appid"`
	MchId          string   `xml:"mchid"`
	NonceStr       string   `xml:"nonce_str"`
	OpenId         string   `xml:"openid"`
	PartnerTradeNo string   `xml:"partner_trade_no"`
	ReUserName     string   `xml:"re_user_name"`
	Sign           string   `xml:"sign"`
	SpbillCreateIp string   `xml:"spbill_create_ip"`
}

//转账响应
type PostResponse struct {
	XMLName        xml.Name `xml:"xml"`
	ReturnCode     string   `xml:"return_code"`
	ReturnMsg      string   `xml:"return_msg"`
	MchAppId       string   `xml:"mch_appid"`
	MchId          string   `xml:"mchid"`
	DeviceInfo     string   `xml:"device_info"`
	NonceStr       string   `xml:"nonce_str"`
	ResultCode     string   `xml:"result_code"`
	RrrCode        string   `xml:"err_code"`
	RrrCodeDes     string   `xml:"err_code_des"`
	PartnerTradeNo string   `xml:"partner_trade_no"`
	PaymentNo      string   `xml:"payment_no"`
	PaymentTime    string   `xml:"payment_time"`
}

//查询请求
type QueryRequest struct {
}

//查询相应
type QueryResponse struct {
}

type Transfer struct {
	Api           string
	Pem           Pem
	Request       *PostRequest
	Response      *PostResponse
	QueryRequest  *QueryRequest
	QueryResponse *QueryResponse
}

func (handler *Transfer) Post(pem Pem, requ *PostRequest, resp *PostResponse) bool {
	handler.Api = API
	handler.Request = requ
	handler.Response = resp
	var signParam []map[string]string
	amount := strconv.Itoa(handler.Request.Amount)
	signParam = append(signParam, map[string]string{"key": "amount", "val": amount})
	signParam = append(signParam, map[string]string{"key": "check_name", "val": handler.Request.CheckName})
	signParam = append(signParam, map[string]string{"key": "desc", "val": handler.Request.Desc})
	signParam = append(signParam, map[string]string{"key": "device_info", "val": handler.Request.DeviceInfo})
	signParam = append(signParam, map[string]string{"key": "mch_appid", "val": handler.Request.MchAppId})
	signParam = append(signParam, map[string]string{"key": "mchid", "val": handler.Request.MchId})
	signParam = append(signParam, map[string]string{"key": "nonce_str", "val": handler.Request.NonceStr})
	signParam = append(signParam, map[string]string{"key": "openid", "val": handler.Request.OpenId})
	signParam = append(signParam, map[string]string{"key": "partner_trade_no", "val": handler.Request.PartnerTradeNo})
	signParam = append(signParam, map[string]string{"key": "re_user_name", "val": handler.Request.ReUserName})
	signParam = append(signParam, map[string]string{"key": "spbill_create_ip", "val": handler.Request.SpbillCreateIp})
	signParam = append(signParam, map[string]string{"key": "key", "val": handler.Request.Key})
	handler.Request.Sign = generateSign(signParam)
	data, _ := xml.Marshal(handler.Request)
	tls := loadPEM(pem.Cert, pem.Key, pem.Ca)
	tr := &http.Transport{
		TLSClientConfig: tls,
	}
	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest("POST", handler.Api, strings.NewReader(string(data)))
	req.Header.Set("Accept", "text/xml,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Charset", "UTF-8")
	res, err := client.Do(req)
	if err != nil {
		fmt.Sprintln(err)
		panic(err)
		return false
	}
	defer res.Body.Close()
	content, rerr := ioutil.ReadAll(res.Body)
	if rerr != nil {
		return false
	}
	xml.Unmarshal([]byte(content), &handler.Response)
	return true
}

//加载CA证书
func loadPEM(CertFilePath, KeyFilePath, CAFilePath string) *tls.Config {
	cert, err := tls.LoadX509KeyPair(CertFilePath, KeyFilePath)
	if err != nil {
		panic(err)
		return nil
	}
	caData, err := ioutil.ReadFile(CAFilePath)
	if err != nil {
		panic(err)
		return nil
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caData)
	var tlsConfig *tls.Config
	tlsConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      pool,
	}
	return tlsConfig
}

func (handler *Transfer) Query() {

}

//生成签名
func generateSign(signParam []map[string]string) string {
	var paramStr string
	l := len(signParam)
	for i := 0; i < l; i++ {
		if signParam[i]["val"] == "" {
			continue
		}
		if i == (l - 1) {
			paramStr = paramStr + signParam[i]["key"] + "=" + signParam[i]["val"]
		} else {
			paramStr = paramStr + signParam[i]["key"] + "=" + signParam[i]["val"] + "&"
		}
	}
	signChar := md5.Sum([]byte(paramStr))
	sign := fmt.Sprintf("%x", signChar)
	return sign
}
