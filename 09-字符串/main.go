

package main 

import (
	"fmt"
	"unicode/utf8"
)

func main(){

	s := "你好golang"
	fmt.Println(s,len(s))
	// len - 12 字节数

	for _, b:= range []byte(s) {
		fmt.Printf("%X\n",b)
	}
	
	for i, ch := range s {
		// ch -- rune -- int32 -- 四字节的整数
		fmt.Printf("(%d %X)\n",i,ch)
	}

	// 字符数 8 
	fmt.Println("rune count:",utf8.RuneCountInString(s))
	// rune count: 8

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c \n",ch)
	}

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)",i,ch)
	}
}