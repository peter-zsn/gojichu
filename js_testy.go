package main

import (
	"fmt"
	"sort"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	//"strconv"
	//"time"
	"encoding/hex"
)

// 构造请求签名sig
func createSig(params map[string]string) string {
	str := ""
	keys := []string{}
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		str += fmt.Sprintf("%s=%s", k, params[k])
	}
	fmt.Println(str)
	key := "7de1a69814ce94eff83bb3ae9dde3ad5"
	//APP_KEY_DeBase64, err := base64.StdEncoding.DecodeString(key)
	//if err != nil {
	//	log.Fatal("Invaild APP_KEY")
	//}
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(str))
	fmt.Println(hex.EncodeToString(mac.Sum(nil)))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func main() {
	//args := map[string]string{
	//	"access_token": "L0Z0FwhhZy9tqBVS7KusByDr7cFGtkvMf3cX5WwWBMwsLDf29mY09w==",
	//	"open_id":      "5e2vJlH3FYmjGmO4j4+ty9sW69VZXDFrgOsqjvwfUKc6EJWhk2WAAA==",
	//	"client_id":    "3667",
	//	"format":       "json",
	//}
	//args["state"] = strconv.FormatInt(time.Now().UnixNano(), 10)
	//args["sig"] = createSig(args)
	//fmt.Println(args)
	param := map[string]string{
		"access_token": "ec9e57913c5b42b282ab7b743559e1b0",
		"formart": "json",
		"oauth_consumer_key":"10000003",
		"openid":"483e74132c9421fca0196d7b94ff0d7a",
	}

	param["sig"] = createSig(param)
	fmt.Println(param["sig"])

}
