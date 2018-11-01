

package main 

import(
	"fmt"
)

func main(){

	// 字典 无序
	// 1.map使用哈希表，必须可以比较相等
	// 2.除了slice，map，function的内建类型都可以作为key
	// 3.struct类型不包含上述字段，也可以作为key
	m := map[string]string{
		"name": "tom",
		"course": "go",
		"site": "video",
	}

	m2 := make(map[string]string)//map[]

	var m3 map[string]int //map[]

	fmt.Println(m,m2,m3)

	// 遍历
	for k, v := range m {
		fmt.Println(k,v)
	}

	courseName := m["course"]
	fmt.Println(courseName)
	// 没有key - "a"，使用不会报错，是默认值
	fmt.Println(m["a"])

	if c, ok := m["names"]; ok {
		fmt.Println(c)
	} else {
		fmt.Println("key names does not exist")
	}

	// 删除m的属性
	delete(m, "name")
	fmt.Println(m["name"])

	// fmt.Println(
	// 	lengthOfNonRepeatingSubStr("abcabcbb"))
	// fmt.Println(
	// 	lengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("pwwkew"))
	// fmt.Println(
	// 	lengthOfNonRepeatingSubStr(""))
	// fmt.Println(
	// 	lengthOfNonRepeatingSubStr("b"))
	// fmt.Println(
	// 	lengthOfNonRepeatingSubStr("abcdef"))
}

// 寻找最长不含有重复字符的子串
func lengthOfNonRepeatingSubStr(s string) int {
	// pwwkew
	lastOccurred := make(map[rune]int)
	fmt.Println("lastOccurred:",lastOccurred)//map[]
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		fmt.Println("i,ch",i,ch)
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastOccurred[ch] + 1
			fmt.Println("start:",start)
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
			fmt.Println("maxLength:",maxLength)
		}
		lastOccurred[ch] = i
	}
	fmt.Println(lastOccurred)
	return maxLength
}






























