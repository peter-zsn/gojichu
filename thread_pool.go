package main

import (
	"fmt"
	"time"
)

type POOL struct {
	Queue	chan func() error			// 任务队列
	RoutineNumber 	int					// 开启的线程数
	QueueSize int						// 任务的长度
	
	Result chan error					// 结果队列长度
	FinashCallback func()				// 任务处理完之后的回调
}


// 初始化pool
func (self *POOL)Init(workNum int, queuesize int){
	self.RoutineNumber = workNum			// 使用的线程数
	self.QueueSize = queuesize				// 任务长度
	
	
	self.Queue = make(chan func()error, queuesize)			//任务队列
	self.Result = make(chan error, queuesize)				// 结果队列
}


// 开始
func (self *POOL)Start(){
	
	for i :=0; i < self.RoutineNumber; i++{			// 开启线程
		go func(){
			for{
				task, ok := <- self.Queue			// 从任务队列读取任务
				if !ok{
					break
				}
				err := task()						// 执行任务
				self.Result <- err					// 结果存放至结果队列
			}
		}()
	}
	
	for j :=0; j < self.QueueSize; j ++{			// 等待所有任务线程执行完毕，读取结果
		res, ok := <- self.Result
		if !ok{
			break
		}
		if res != nil{
			fmt.Println(res)
		}
	}
	
	
	if self.FinashCallback != nil{					// 结果完成的回调
		self.FinashCallback()
	}
	self.Stop()
}

func (self *POOL)Stop(){
	close(self.Queue)
	close(self.Result)
}


// 添加任务
func (self *POOL)AddTask(task func() error){
	self.Queue <- task
}


// 设置回调函数
func (self *POOL)SetFinashCallback(f func()){
	self.FinashCallback = f
}


func main(){
	var pool POOL
	
	urls := []string{"葛地主的网址", "宋锤子的网址", "梁胖子的网址", "张少帅的网址"}
	
	pool.Init(3, len(urls))
	for _, url := range urls{
		a := url
		pool.AddTask(func() error{
			return ShowUrl(a)
		})
	}
	
	pool.SetFinashCallback(ShowFinish)
	pool.Start()
}


func ShowUrl(url string)error{
	time.Sleep(1*time.Second)
	fmt.Println(url)
	return nil
}

func ShowFinish(){
	fmt.Println("job is over")
}