// access_token.go
package access_token

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	API_GET_ACCESS_TOKEN = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&"
	TOKEN_FILE_FOLDER    = "access_token/at"
)

type Response struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expire_in"`
}

type AccessToken struct {
	Api            string
	AppId          string
	Secret         string
	Response       Response
	IsCache        bool
	CacheDirectory string
}

func (handler *AccessToken) Get() bool {
	filename := ""
	if handler.IsCache {
		filename = generateFilename(handler.AppId)
		res, existReponse := readFile(handler.CacheDirectory, filename)
		if res {
			handler.Response = existReponse
			return res
		} else {
			reqRes := handler.getFromServe()
			if reqRes {
				handler.Response.ExpiresIn = time.Now().Unix() + 2400
				saveFile(handler.CacheDirectory, filename, handler.Response)
			}
			return reqRes
		}
	} else {
		return handler.getFromServe()
	}

	//	res, existReponse := readFile(handler.CacheDirectory, filename)
	//	if res {
	//		handler.Response = existReponse
	//		return res
	//	} else {
	//		handler.Api = API_GET_ACCESS_TOKEN
	//		params := "appid=" + handler.AppId + "&secret=" + handler.Secret
	//		resp, err := http.Get(handler.Api + params)
	//		if err != nil {
	//			panic(err)
	//		}

	//		defer resp.Body.Close()
	//		body, err := ioutil.ReadAll(resp.Body)
	//		if err != nil {
	//			panic(err)
	//		}
	//		enErr := json.Unmarshal(body, &handler.Response)
	//		if enErr != nil {
	//			panic(enErr)
	//		}
	//		if handler.IsCache {
	//			handler.Response.ExpiresIn = time.Now().Unix() + 2400
	//			saveFile(handler.CacheDirectory, filename, handler.Response)
	//		}
	//		if handler.Response.AccessToken != "" {
	//			return true
	//		} else {
	//			return false
	//		}
	//	}
}

func (handler *AccessToken) getFromServe() bool {
	handler.Api = API_GET_ACCESS_TOKEN
	params := "appid=" + handler.AppId + "&secret=" + handler.Secret
	resp, err := http.Get(handler.Api + params)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	enErr := json.Unmarshal(body, &handler.Response)
	if enErr != nil {
		panic(enErr)
	}
	if handler.Response.AccessToken != "" {
		return true
	} else {
		return false
	}
}

func generateFilename(key string) string {
	return key + ".at"
}

func readFile(saveDirecoty string, filename string) (bool, Response) {
	var res Response
	filePath := saveDirecoty + "/" + filename
	fileInfo, fierr := os.Stat(filePath)
	if fierr != nil {
		return false, res
	}
	fp, ferr := os.Open(filePath)
	if ferr != nil {
		panic(ferr)
	}
	defer fp.Close()
	content := make([]byte, fileInfo.Size())
	readByte, rerr := fp.Read(content)
	if readByte == 0 || rerr != nil {
		panic(rerr)
	}
	merr := json.Unmarshal(content, &res)
	if merr != nil {
		return false, res
	} else {
		if res.ExpiresIn < time.Now().Unix() {
			return false, res
		} else {
			return true, res
		}
	}
}

func saveFile(saveDirecoty string, filename string, res Response) bool {
	os.MkdirAll(saveDirecoty, 0755)
	filePath := saveDirecoty + "/" + filename
	fp, ferr := os.Create(filePath)
	defer fp.Close()
	if ferr != nil {
		panic(ferr)
	}
	bytes, merr := json.Marshal(res)
	if merr != nil {
		panic(merr)
	}
	writeByte, werr := fp.Write(bytes)
	if writeByte != 0 && werr == nil {
		return true
	} else {
		return false
	}
}
