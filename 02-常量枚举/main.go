

package main 


import (
	"fmt"
	"math/cmplx"
	"math"
)


func euler(){
	// 复数
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))

	fmt.Println(cmplx.Exp(1i*math.Pi)+1)
	//(0+1.2246467991473515e-16i)

	fmt.Printf("%.3f\n",cmplx.Exp(1i*math.Pi)+1)
	// (0.000+0.000i)
}

func tringle(){
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a,b))
}
func calcTriangle(a, b int) int {
	var c int
	// 类型强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts() {
	// 常量
	const (
		filename = "abc"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}


func enums() {
	// 枚举
	const (
		cpp = 0
		java = 1
		python = 2
		golang = 3
	)
	fmt.Println(cpp, java, python, golang)

	const (
		// 自增类型
		one = iota
		_
		three
		four
		five
	)
	fmt.Println(one, three, four, five)

	const (
		// 向左位移
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b,kb,mb,gb,tb,pb)
}

func main(){
	euler()
	tringle()
	consts()
	enums()
}














