

package main 

import "fmt"

// 全局变量
var cc = 33
// 集中定义变量
var (
	aa = 3
	ss = "ss"
	bb = true
)


func variableZeroValue(){
	// 不赋值 默认值
	var a int //0
	var s string//空字符串
	fmt.Printf("%d %q\n",a, s)//%q可以打印出空字符串
}

func variableInitialValue(){
	// 声明变量 并且赋初始值
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a,b,s)
}

func variableTypeDeduction(){
	// 可以省略类型赋值 会自行判断类型
	var a, b, c , s = 3, 4, true, "def"
	fmt.Println(a,b,c,s)
}

func variableShorter() {
	// 简短声明赋值方式
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a,b,c,s)
}

func main(){
	fmt.Println("hello go")

	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
}
















