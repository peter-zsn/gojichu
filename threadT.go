package main

import (
	"strings"
	"io/ioutil"
	"encoding/json"
	"net/url"
	"net/http"
	"time"
	"strconv"
	"fmt"
)

var (
	httpClient = http.Client{
		Timeout: 5 * time.Second,
	}
)

func main(){
	var i int
	channel := make(chan map[string]string, 1000)
	for i = 1; i < 50; i++ {
		go func(id int){
			channel <- testfun(id)
		}(i)
	}
	var j int
	for j = 1; j < 50; j++{
		out := <- channel
		if(out["response"] == "fail"){
			fmt.Println(out["id"])
		}
	}
}

func testfun(i int)map[string]string{
	url := "http://192.168.7.103:8006/account/getuser"
	args := map[string]string{}
	args["user_id"] = strconv.Itoa(i)
	out := map[string]string{}
	HttpPost(url, args, &out)
	out["id"] = strconv.Itoa(i)
	return out
}

func SetBaseParams(params map[string]string) string{
	kv := []string{}
	for k, v := range params{
		kv = append(kv, k+"="+url.QueryEscape(v))
	}
	return strings.Join(kv, "&")
}

// post请求
func HttpPost(url string, params map[string]string, out interface{})error{
	args := SetBaseParams(params)
	body := strings.NewReader(args)
	resp, err := httpClient.Post(url, "application/x-www-form-urlencoded", body)
	if err != nil{
		return nil
	}
	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return err
	}
	err = json.Unmarshal(resBody, out)
	if err != nil{
		return  err
	}
	return  nil
}