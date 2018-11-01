

package main  

import(
	"fmt"
)

func main(){

	arr := [...]int{1,2,3,4,5,6,7,8,9,10}

	//slice的四种方式
	fmt.Println("arr[2:6]=",arr[2:6])
	fmt.Println("arr[:6]=",arr[:6])
	fmt.Println("arr[2:]=",arr[2:])
	fmt.Println("arr[:]=",arr[:])

	s1 := arr[2:6]
	fmt.Println("s1=",s1)//s1= [3 4 5 6]

	updateSlice(s1)
	fmt.Println("更新后：s1=",s1)//s1= [100 4 5 6]
	fmt.Println("arr=",arr)

	printArray(arr[:])
	fmt.Println("arr=",arr)
	// Slice本身没有数据，是对底层array的一个view

	// slice可以继续slice，也就是reslice
	s1 = s1[1:]
	fmt.Println("reslice:s1=",s1)//[4 5 6]

	// 直接取s1[3]报错，因为超出index边界
	// fmt.Println(s1[3])
	//panic: runtime error: index out of range

	// 并不报错，而且能取出值
	fmt.Println(s1[2:4])//[6 7]

	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",s1,len(s1),cap(s1))
	// s1=[4 5 6], len(s1)=3, cap(s1)=7

	fmt.Println(s1[2:7])//[6 7 8 9 10]

	// 超出了cap容量报错
	// fmt.Println(s1[2:8])
	// panic: runtime error: slice bounds out of range
	s2 := append(s1,11)
	s3 := append(s2,22)
	s4 := append(s3,33)
	s5 := append(s4,44)
	fmt.Println(s2,s3,s4,s5)
	fmt.Println("append后：arr=",arr)
	// append后：arr= [111 2 100 4 5 6 11 22 33 44]
	// 注：在不超出cap的情况下，就是爱改变原来的数组arr，所以操作slice就是在操作数组

	s6 := append(s5,55)
	fmt.Println(s6)
	fmt.Println("append后：arr=",arr)
	// append后：arr= [111 2 100 4 5 6 11 22 33 44]
	// 超出cap后，程序会在内存中重新开辟空间，创建一个cap更大的数组，把原来arr的值全部拷贝过去
	// 这个新数组我们是不可见的，程序自动创建的

	sliceOther()
}

func updateSlice(s []int) {
	s[0] = 100
}
// 参数为slice
func printArray(s []int){
	s[0] = 111
	for i, v := range s {
		fmt.Println(i,v)
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
}

func sliceOther(){

	var s []int//slice的默认值是nil

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}

	// 声明并赋默认值
	s1 := []int{1,2,3,4}
	fmt.Println(s1)

	// 创建并开辟空间 16-len 32-cap
	s2 := make([]int, 16, 32)
	fmt.Println(s2)//[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]

	copy(s2,s1)
	printSlice(s2)
	fmt.Println("copy后：",s2)

	s2 = append(s2[:3], s2[4:]...)
	fmt.Println("删除4：",s2)//[1 2 3 0 0 0 0 0 0 0 0 0 0 0 0]
}













