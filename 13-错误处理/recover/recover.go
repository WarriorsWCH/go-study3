package main 

import "fmt" 

func tryRecover() {
	defer func() {
		// panic：不要经常使用
		// 停止当前函数执行
		// 一直向上返回，执行每一层的defer
		// 如果没有遇见recover，程序退出
		// recover:
		// 仅在defer调用中使用
		// 获取panic的值
		// 如果无法处理，可重新panic
		r := recover()

		if r == nil {
			fmt.Println("Nothing to recover")
			return
		}
		// 判断错误是不是error
		if err, ok := r.(error); ok {
			fmt.Println("error occurred:", err)
		} else {
			// 不是error 直接panic
			panic(fmt.Sprintf("i don`t konw what to do:%v ",r))
		}
	}()

	// panic(errors.New("this is an error"))

	// b := 0
	// a := 5 / b
	// fmt.Println(a)
	//error occurred: runtime error: integer divide by zero

	// panic(123)//panic: i don`t konw what to do:123
}

func main(){
	tryRecover()
}