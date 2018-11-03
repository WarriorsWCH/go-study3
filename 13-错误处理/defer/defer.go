

package main 

import (
	"fmt"
	"os"
	"bufio"
	"groot/imooc/12-函数式编程/fib"
)

func tryDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		// 不管是return 还是panic defer都会被调用
		if i == 30 {
			fmt.Println("it is 30")
			// return
			// panic("111")
		}
	}
}


func wrteFile(filename string) {
	// 0666是权限
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE|os.O_WRONLY,0666)

	// 错误处理
	if err != nil {
		// 错误的类型是*PathError
		// !ok 如果不是，那就不知道是什么了，直接panic
		if pathError, ok := err.(*os.PathError); !ok{
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
			// open, fib.txt, file exists
		}
		return
	}
	// 关闭文件
	defer file.Close()

	// 直接用file写文件会比较慢，所以使用bufio，
	// 带buffer的，先写到内存中，当内存中达到一定程度再全部导过去
	writer := bufio.NewWriter(file)
	// 把内存中的数据导入文件
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main () {
	tryDefer()
	wrteFile("fib.txt")
}



















