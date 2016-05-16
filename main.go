// main.go
package main

import (
	"fmt"
	_ "time"
	"weixin-api/access_token"
	_ "weixin-api/message"
)

const (
	API_CERT_FILE = "pem/apiclient_cert.pem"
	API_KEY_FILE  = "pem/apiclient_key.pem"
	API_CA_FILE   = "pem/rootca.pem"
)

func main() {
	//	handler := new(transfer.Transfer)
	//	var req transfer.PostRequest
	//	req.MchAppId = "wx385749384926ddeb"
	//	req.MchId = "1252919701"
	//	req.DeviceInfo = "xzd_device"
	//	req.NonceStr = "abcewewfewf"
	//	req.PartnerTradeNo = "20160501113001"
	//	req.OpenId = "o9Rbts6sAS34xNXz_7tHyKiKKrxU"
	//	req.CheckName = transfer.CHECK_NAME_OPTION_NO
	//	req.ReUserName = ""
	//	req.Amount = 100
	//	req.Desc = "22222test transfer"
	//	req.SpbillCreateIp = "127.0.0.1"
	//	req.Key = "30b3efbf3656ba263dbb505c5d9bd1dc"
	//	var res transfer.PostResponse
	//	var pem transfer.Pem
	//	pem.Cert = API_CERT_FILE
	//	pem.Key = API_KEY_FILE
	//	pem.Ca = API_CA_FILE
	//	handler.Post(pem, req, res)
	//	fmt.Printf("%+v", handler.Request)
	//	fmt.Println("\n")
	//	fmt.Printf("%+v", handler.Response)
	handler := new(access_token.AccessToken)
	handler.IsCache = true
	handler.CacheDirectory = "cache"
	handler.AppId = "wx385749384926ddeb"
	handler.Secret = "be51e251b884156db5df7cc92807d7aa"
	handler.Get()
	fmt.Printf("%+v", handler.Response)
	//	if res {
	//		messageHanlder := new(message.WxSend)
	//		messageHanlder.AccessToekn = handler.Response.AccessToken
	//		sendReq := new(message.SendRequest)
	//		sendReq.SetParam("first", "尊敬的业主您好！荟当家平台本次补贴您的物业费XX元")
	//		sendReq.SetParam("tradeDateTime", time.Now().Format("2006-01-02 15:04:05"))
	//		sendReq.SetParam("tradeType", "消费抵扣物业费")
	//		sendReq.SetParam("curAmount", "129")
	//		sendReq.SetParam("remark", "我们是会当家")
	//		sendReq.TemplateId = "U-ncYr9CTpvbRbmND23KEdgvX4cms3Eoo3yV3rRKs-w"
	//		sendReq.Touser = "o9Rbts6sAS34xNXz_7tHyKiKKrxU"
	//		sendReq.Url = "http://hdj.sunyeart.com"
	//		messageHanlder.Send(sendReq)
	//		fmt.Println(messageHanlder.Response)
	//	} else {
	//		fmt.Println("get access_token fail")
	//	}

}
