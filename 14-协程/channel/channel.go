package main 

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for n := range c{
		// n := <- c
		// fmt.Println(n)
		fmt.Printf("Worker %d received %c\n",id, n)
	}
}

func createWorker(id int) chan<-int {
	c := make(chan int)
	go worker(id, c)
	return c
}
func chanDemo() {
	// var c chan int // c == nil
	// c := make(chan int)
	// go func() {
	// 	for{
	// 		n := <- c
	// 		fmt.Println(n)
	// 	}
	// }()
	// go worker(0, c)

	// chan<- int只能发数据
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		// channels[i] = make(chan int)
		// go worker(i, channels[i])
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++{
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++{
		channels[i] <- 'A' + i
	}
	// c <- 1
	// c <- 2
	// 读数据需要再开一个协程，否则报错
	// n := <- c//fatal error: all goroutines are asleep - deadlock!
	// fmt.Println(n)
	time.Sleep(time.Millisecond)
}


func main() {
	// chanDemo()
	// bufferedChannel()
	channelClose()
}
	

func bufferedChannel() {
	// (chan int, 3)缓冲channel，可以提升性能
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	// 发送方close
	// 接收方要关闭for n := range c
	close(c)
	time.Sleep(time.Millisecond)
}
// 不要通过共享内存来通信（比如一个flag的false和true），通过通信来共享内存


















