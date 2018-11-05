

package main 


import(
	"fmt"
	"time"
	"math/rand"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(
				time.Duration(rand.Intn(1500)) *
					time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n",
			id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}


func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)

	var values []int
	// 返回值是<-chan Time 多少秒后
	tm := time.After(10 * time.Second)
	// // 返回值是<-chan Time 每隔段时间送个值过来
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan <- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
			// 会等段时间才会case c1 和 c2 ，会先打印很多default
			case n := <-c1:
				values = append(values, n)
				fmt.Println("c1")
			case n := <-c2:
				values = append(values, n)
				fmt.Println("c2")
			case activeWorker <- activeValue:
				// 把第一个拿掉也就是values[0]
				values = values[1:]
			case <-tick:
				fmt.Println("queue len =", len(values))

			case <-time.After(800 * time.Millisecond):
				// 每次select的的时间
				fmt.Println("timeout")
			case <-tm:
				fmt.Println(len(values),"bye")
				return
			// default:
			// 	fmt.Println("no values received")
		}
	}
}
















