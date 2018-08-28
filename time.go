package main

import (
	"fmt"
	"time"
	//"strconv"
	//"reflect"
)

func main() {
	//fmt.Println(time.Second)
	a := time.Now().Unix()
	fmt.Println(a)
	//scheduleTime := "1504080559"
	//scheduleTime = fmt.Sprintf("%s", []byte(scheduleTime)[:10])
	//fmt.Println(scheduleTime)
	//fmt.Println(reflect.TypeOf(scheduleTime))
	//send_time, err := strconv.ParseInt(scheduleTime, 10, 64)
	//fmt.Println(err)
	//scheduleTime = time.Unix(send_time, 0).Format("2006-01-02 15:04:05")
	//fmt.Println(scheduleTime)
	//fmt.Println(reflect.TypeOf(scheduleTime)).

	//scheduleTime :="1504080559"	// 发送时间
	//
	//if scheduleTime != ""{				// 定时发短信
	//	scheduleTime = fmt.Sprintf("%s", []byte(scheduleTime)[:10])
	//	send_time, err := strconv.ParseInt(scheduleTime, 10, 64)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	scheduleTime = time.Unix(send_time, 0).Format("2006-01-02 15:04:05")
	//}
	//fmt.Println(scheduleTime)
}
