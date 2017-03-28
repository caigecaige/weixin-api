// message.go
package message

import (
	"crypto/tls"
	"encoding/json"
	_ "fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	SEND_API  = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token="
	SEND_SUCC = "ok"
)

type SendRequest struct {
	Touser     string                       `json:"touser"`
	TemplateId string                       `json:"template_id"`
	Url        string                       `json:"url"`
	Data       map[string]map[string]string `json:"data"`
}

func (handler *SendRequest) SetParam(key string, val string) {
	if len(handler.Data) == 0 {
		handler.Data = make(map[string]map[string]string)
	}
	param := make(map[string]string)
	param["value"] = val
	handler.Data[key] = param
}

type SendResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	MsgId   int64  `json:"msgid"`
}

type WxSend struct {
	Api         string
	AccessToekn string
	Request     *SendRequest
	Response    *SendResponse
}

func (handler *WxSend) Send(sreq *SendRequest) bool {
	//fmt.Printf("%+v", sreq)
	handler.Request = sreq
	handler.Api = SEND_API + handler.AccessToekn
	data, _ := json.Marshal(handler.Request)
	//fmt.Printf("%s", data)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //不检验证书
	}
	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest("POST", handler.Api, strings.NewReader(string(data)))
	req.Header.Set("Accept", "text/xml,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Charset", "UTF-8")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content, rerr := ioutil.ReadAll(res.Body)
	if rerr != nil {
		return false
	}
	json.Unmarshal([]byte(content), handler.Response)
	return true
}
