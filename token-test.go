package main

import (
	"fmt"
	"encoding/base64"
	"strings"
	"time"
	"strconv"
	"errors"
	"crypto/md5"
	"io"
	"math"
)

var KEY = 240897
var KEYSTR = "240897"
var SESSION_COOKIE_AGE = 864000

// 解密token，得到userId和过期时间的时间戳， 如果过期或者比对不上返回错误信息和空
func decode(token string)(error, string, string){
	if token == ""{
		return nil, "", ""
	}
	tokenLen := 4 - len(token) % 4
	for i:= 0; i < tokenLen; i++{
		token += "="
	}
	tokenDed, err := base64.StdEncoding.DecodeString(token)
	if err != nil{
		return err, "", ""
	}
	deStr := strxor(tokenDed)
	a:= strings.Split(deStr, "|")
	if len(a) != 3{
		return errors.New("strxor error"), "", ""
	}
	userId, exTime, code := a[0], a[1], a[2]
	return nil, userId, exTime
	now := time.Now().Unix()
	exTimeInt, _ := strconv.ParseInt(exTime, 10, 64)
	if now > exTimeInt{
		return errors.New("time out"), "", ""
	}
	relCode := createCheckCode(userId, exTimeInt)
	if code != relCode{
		return errors.New("code error"), "", ""
	}
	return nil, userId, exTime
}

// 对字节数组进行亦或运算
func strxor(s []uint8)string{
	newKey := KEY & 0xff
	for i, c := range s{
		d := c ^ uint8(newKey)
		s[i] = d
	}
	return string(s)
}

// 根据userid和过期时间得到验证码------异或运算
func createCheckCode(userId string, exTime int64)string{
	userIdInt, _ := strconv.ParseInt(userId, 10, 64)
	key , _ := strconv.ParseInt(KEYSTR, 10, 64)
	c := userIdInt ^ exTime ^ key
	return strconv.FormatInt(c,10)
}

// 根据userId进行加密得到token
func createToken(userId string)string{
	now := time.Now().Unix()
	exTim := now + int64(SESSION_COOKIE_AGE)
	code := createCheckCode(userId, exTim)
	token := fmt.Sprintf("%s|%s|%s" , userId, strconv.FormatInt(exTim, 10), code)
	tokenArry := []byte(token)
	TwoToken := strxor(tokenArry)
	relToken := base64.StdEncoding.EncodeToString([]byte(TwoToken))
	relToken = strings.Replace(relToken, "=", "", -1)
	return relToken
}

func get_token() string {
	m5 := md5.New()
	youxue_key  := "7c7edf96a6a5a3c29040cf9fa87cb7a8"
	m5Str := time.Now().Format("2006-01-02 15:04") + youxue_key
	io.WriteString(m5, m5Str)
	return fmt.Sprintf("%x", m5.Sum(nil))
}

func main()  {
	token := "MH0wNDI0MTE1NDA5fTA0MjU4Njc3MDU"
	err, userId, exTime := decode(token)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(math.MaxInt8)
	fmt.Println(userId, exTime)
	//3618  34704
	newToken := createToken("7694506")
	fmt.Println(newToken)
	fmt.Println(get_token())
}