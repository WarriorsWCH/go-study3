
package main 

import (
	"fmt"
	"bufio"
	"io"
	"strings"
	"groot/imooc/12-函数式编程/fib"
)

// 闭包
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder{
	return func(v int)(int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main(){
	// 函数时一等公民：参数，变量，返回值都可以是函数
	// 高阶函数：参数也可以是函数
	
	//正统函数式编程
	// 不可变性：不能有状态，只有常量和函数
	// 函数只能有一个参数

	f := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0+1+.....+%d\n",f(i))
	}

	a := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n",
			i, s)
	}

	var fb intGen = fib.Fibonacci()
	printFileContents(fb)
}

// 为函数实现接口(给fibonacci实现接口)
type intGen func() int
// 函数作为接受者(g intGen)
// Read接口 系统的接口如下
// type Reader interface {
	//Read(p []byte) (n int, err error)
// }
// 实现这个接口就可以像读文件一样读取数据
func (g intGen) Read(p []byte) (n int, err error) {
	// 获取下一个元素 g() -- fibonacci()
	next := g()
	// 如果结果大于10000 就停止读取
	if next > 10000 {
		return 0, io.EOF
	}
	// 转字符串
	s := fmt.Sprintf("%d\n", next)
	// 把字符串写入p []byte中
	// 利用strings把结果写入p []byte中
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}



















