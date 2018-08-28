package main

import (
	"fmt"
	"runtime"
)

func ThreadFun(){
	for i:=1; i < 3; i++ {
		runtime.Gosched()
		fmt.Println("this is threadfun.......")
	}
}

func main() {
	fmt.Println("this is start")
	go ThreadFun()
	ThreadFun()
	fmt.Println("this is over")

}
