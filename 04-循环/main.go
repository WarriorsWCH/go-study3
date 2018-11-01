
package main  

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

// 转二进制
// 13 -- 1101
// 13对2取模 -->1 余6
// 6对2取模 -->0 余3
// 3对2取模 -->1 余1
// 1对2取模 -->1 余0
// 1011 倒过来 1101 就是13的二进制
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		// 倒过来
		result = strconv.Itoa(lsb) + result
	}
	return result
}


func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main(){
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(12312312),
		convertToBin(0),
	)

	printFile( "./test.txt")
}

// 1.for, if 后面的条件没有括号
// 2.if条件里也可以定义变量
// 3.没有while
// 4.switch不需要break，也可以直接switch多个条件










