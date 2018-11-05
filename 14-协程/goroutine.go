

package main 

import (
	"fmt"
	"time"
	// "runtime"
)

func main(){
	var a [10]int
	for i := 0; i < 10; i++ {
		// 并发执行函数
		go func(i int) {
			for {
				a[i]++
				fmt.Printf("hello from goroutine %d\n",i)
				// runtime.Gosched()//交出控制权
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
// 协程goroutine
// 轻量级“线程”
// 非抢占式多任务处理，由协程主动交出控制权（由协程主动交出控制权）
// 线程是抢占式多任务处理
// 编译器、解释器、虚拟机层面的多任务
// 多个协程可能在一个或多个线程上运行

// goroutine的定义：
// 任何函数只需要加上go就能送给调度器运行
// 不需要再定义时区分是否是异步函数
// 调度器在合适的点进行切换
// 使用-race来检测数据访问冲突

// goroutine可能切换的点
// I/O,select
// channel
// 等待锁（有时）
// 函数调用
// untime.Gosched()

















