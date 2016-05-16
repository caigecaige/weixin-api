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

	handler := new(access_token.AccessToken)
	handler.IsCache = true
	handler.CacheDirectory = "cache"
	handler.AppId = "appid"
	handler.Secret = "appsecrect"
	handler.Get()
	fmt.Printf("%+v", handler.Response)

}
