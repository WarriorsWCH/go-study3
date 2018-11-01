

package main 

import(
	"fmt"
)

func printArray(arr [5]int){
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i,v)
	}
}

func printArray2(arr *[5]int){
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i,v)
	}
}

func main(){
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)

	// 声明
	var arr1 [5]int
	// 简短声明 必须要赋值
	arr2 := [3]int{1,3,5}
	// 不固定长度 需要使用... 否则是slice
	arr3 := [...]int{1,4,6,8,10}

	var grid [4][5]int

	fmt.Println(arr1,arr2,arr3)
	fmt.Println(grid)

	// 数组也是值传递 arr[0] = 100并未改变原来的数组
	printArray(arr3)
	fmt.Println(arr3)//[1 4 6 8 10]
	// printArray(grid)
	// 注意：
	// 1.[10]int和[20]int是不同类型
	// 2.调用func f(arr [10]int) 会拷贝数组 

	// 传递指针
	printArray2(&arr3)
	fmt.Println(arr3)//[100 4 6 8 10]
	// 3.在go语言中一般不直接使用数组，而是切片
}







