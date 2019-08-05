package base

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个goroutine,启动一个任务
	go newThreadTask()
	// main goroutine循环打印
	i := 0
	for {
		i++
		fmt.Println("main goroutine: i = %d\n", i)
		time.Sleep(1 * time.Second) //延时1s
	}
}

func newThreadTask() {
	i := 0
	for {
		i++
		fmt.Println("new goroutine: i = %d\n", i)
		time.Sleep(1 * time.Second) // 延时1s
	}
}
