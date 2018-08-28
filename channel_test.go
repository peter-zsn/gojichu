package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 100)
	go func() {
		for{
			j, more:= <-jobs
			if more{
				fmt.Println("receive job", j)
			}
			else{
				fmt.Println("receive all job")
				return
			}
		}
	}()


}
