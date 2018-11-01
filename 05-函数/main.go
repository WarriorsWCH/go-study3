

package main 

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)
// 返回值类型卸载最后面
func eval(a, b int, op string) (int, error) {
	switch op {
		case "+":
			return a + b, nil
		case "-":
			return a - b, nil
		case "*":
			return a * b, nil
		case "/":
			q, _ := div(a,b)
			return q, nil
		default:
			return 0, fmt.Errorf("unsupported oprationL %s",op)
	}
}
// 函数返回多个值时可以起名字
// 仅用于非常简单的函数
// 对于调用者而言没有区别
func div(a, b int) (q, r int) {
	return a / b, a % b
}

// 函数作为参数
func apply(op func(int, int) int, a, b int) int {
	// refelct反射  Pointer获取指针
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args" + "(%d, %d)\n",opName, a, b)
	return op(a, b)
}

// 可变参数
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main(){

	fmt.Println(eval(3, 4, "*"))
	q, r := div(13, 3)
	fmt.Println(q,r)


	fmt.Println("pow(3,4) is:", apply(
		func(a int, b int) int {
			return int(math.Pow(
				float64(a), float64(b)))
		}, 3, 4))

	fmt.Println("1+2+...+5=", sum(1,2,3,4,5))
}















