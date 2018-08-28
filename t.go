package main

import (
	"fmt"
	"strings"
	"time"
	"crypto/md5"
	"encoding/hex"
)

var KEY  = "chsy1234"

func main()  {
	a := "http://file.m.xueceping.cn/upload_media/knowledge/video/2013/11/08/20131108174410794323.mp4"
	b := strings.Split(a, "/upload_media/")
	fmt.Println(b[0])
	fmt.Println(b[1])
	fmt.Println(len(b))
	if len(b) < 2{
		return
	}
	relData := "/upload_media/" + b[1]
	fmt.Println(relData)
	timeStr := time.Now().Format("200601021504")
	fmt.Println(timeStr)
	needData := KEY + timeStr + relData
	fmt.Println(needData)
	h := md5.New()
	h.Write([]byte(needData))
	chpherStr := hex.EncodeToString(h.Sum(nil))
	result := b[0] + "/" + timeStr + "/" + chpherStr + relData
	fmt.Println(result)
}